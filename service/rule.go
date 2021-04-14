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
	once := strings.TrimSpace(req.OnceLimit)
	day := strings.TrimSpace(req.DayLimit)
	rule := fmt.Sprintf("%s,%s", once, day)

	foodRule := model.MgoRule{
		Title:    title,
		Rule:     rule,
		RuleType: model.RuleTypeFood,
	}

	repository.MgoRuleRepo.Create(foodRule)

	resp := &RuleFoodCreateResponse{}
	resp.Code = 0
	resp.Msg = "SUCCESS"

	//resp.Data.AccessToken = req.AccountNo
	//resp.Data.RefreshToken = req.Password

	return resp, nil
}
