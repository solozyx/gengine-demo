package handler

import "github.com/gin-gonic/gin"

var (
	MineHandler = &mineHandler{}
)

// ------------------------------------------------------------------------

//go:generate go run ../../tool/gen_handler/gen_handler.go

type IMineHandler interface {
	Login() gin.HandlerFunc
}
