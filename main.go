package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gengine/common"
	"gengine/gateway"
	"gengine/repository"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	common.Cfg = new(common.Config)
	if err := viper.Unmarshal(common.Cfg); err != nil {
		panic(err)
	}

	common.InitLog(os.Stdout)
	logrus.Infof("config=%+v", common.Cfg)
	repository.Init(common.Cfg)
	// service.Init(common.Cfg)
}

func main() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	// DEBUG
	//g.Use(gin.LoggerWithWriter(&myGinLogger{}))

	gateway.BindRouter(g) //setContext(),
	//middleware.Options,
	//middleware.SignCheck(),
	//middleware.AuthCheck(),

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}

func setContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("request_uri", ctx.Request.RequestURI)
		ctx.Set("config", common.Cfg)
		ctx.Set("ip", ctx.ClientIP())
		ctx.Next()
	}
}
