package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"gengine/common"
	. "gengine/gateway/model"
	"gengine/repository"
	"gengine/repository/model"
	"github.com/bilibili/gengine/builder"
	genginectx "github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

var (
	ruleSrv *ruleService
)

type IRuleService interface {
	FoodCreate(req RuleFoodCreateRequest) (*RuleFoodCreateResponse, error)
	FoodCheck(req RuleFoodCheckRequest) (*RuleFoodCheckResponse, error)
}

func NewRuleService(ctx context.Context) IRuleService {
	userId, ok := ctx.Value(common.UserIdKey).(int)

	if ruleSrv != nil {
		if ok {
			ruleSrv.UserId = userId
		}
		return ruleSrv
	}

	ruleSrv = &ruleService{}
	if ok {
		ruleSrv.UserId = userId
	}
	return ruleSrv
}

type ruleService struct {
	UserId int
}

func (s *ruleService) FoodCreate(req RuleFoodCreateRequest) (*RuleFoodCreateResponse, error) {
	title := strings.TrimSpace(req.Title)
	// once := strings.TrimSpace(req.OnceLimit)
	// day := strings.TrimSpace(req.DayLimit)
	// rule := fmt.Sprintf("%s,%s", once, day)

	// rule := fmt.Sprintf("%.2f,%.2f", req.OnceLimit, req.DayLimit)

	var rule_else_if_test = `
rule "%s" "%s"
begin

//单次用餐金额需低于 元
onceLimit = %.2f

//单日用餐金额需低于 元
dayLimit = %.2f

once = %s
if once > onceLimit {
	println("超过单次用餐金额限制")
}

day = %s
if day > dayLimit {
	println("超过单日用餐金额限制")
}

end
`

	rule := fmt.Sprintf(rule_else_if_test, title, title,
		req.OnceLimit, req.DayLimit, "%.2f", "%.2f")

	foodRule := model.MgoRule{
		Title:    title,
		Rule:     rule,
		RuleType: model.RuleTypeFood,
	}

	err := repository.MgoRuleRepo.Create(foodRule)
	if err != nil {
		return nil, err
	}

	resp := &RuleFoodCreateResponse{}
	resp.Code = 0
	resp.Msg = "SUCCESS"

	return resp, nil
}

func (s *ruleService) FoodCheck(req RuleFoodCheckRequest) (*RuleFoodCheckResponse, error) {
	res, err := repository.MgoRuleRepo.GetLatestFoodRule()
	if err != nil {
		return nil, err
	}

	// gengine
	dataContext := genginectx.NewDataContext()
	//dataContext.Add("println", fmt.Println)
	dataContext.Add("println", logrus.Println)
	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//resolve rules from string
	once := req.Once
	day := req.Day
	rule := res.Rule
	logrus.Debugf("rule=%+v", rule)

	if once == 0 && day == 0 {
		rule = fmt.Sprintf(rule, 0, 0)
	}

	if once > 0 && day == 0 {
		rule = fmt.Sprintf(rule, once, 0)
	}

	if once == 0 && day > 0 {
		rule = fmt.Sprintf(rule, 0, day)
	}

	if once > 0 && day > 0 {
		rule = fmt.Sprintf(rule, once, day)
	}

	logrus.Debugf("rule=%+v", rule)

	err = ruleBuilder.BuildRuleFromString(rule)
	if err != nil {
		return nil, err
	}

	eng := engine.NewGengine()

	err = eng.Execute(ruleBuilder, true)
	if err != nil {
		return nil, err
	}

	resp := &RuleFoodCheckResponse{}
	resp.Code = 0
	resp.Msg = "SUCCESS"
	resp.Data.Permitted = true
	return resp, nil
}
