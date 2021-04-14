package service

import (
	"context"
	"fmt"
	"gengine/repository"
	"gengine/repository/model"
	"strings"

	"gengine/common"
	. "gengine/gateway/model"
)

var (
	ruleSrv *ruleService
)

type IRuleService interface {
	Create(req RuleFoodCreateRequest) (*RuleFoodCreateResponse, error)
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

func (s *ruleService) Create(req RuleFoodCreateRequest) (*RuleFoodCreateResponse, error) {
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
