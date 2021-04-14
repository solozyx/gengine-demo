package handler

import "github.com/gin-gonic/gin"

var (
	MineHandler = &mineHandler{}
	RuleHandler = &ruleHandler{}
)

// ------------------------------------------------------------------------

//go:generate go run ../../tool/gen_handler/gen_handler.go

type IMineHandler interface {
	Login() gin.HandlerFunc
}

type IRuleHandler interface {
	FoodCreate() gin.HandlerFunc
}
