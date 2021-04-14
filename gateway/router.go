package gateway

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gengine/gateway/handler"
)

func BindRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// 处理静态资源
	//g.Static(common.PathStatic, "./file")

	g.Use(gin.Recovery())

	// 使用Gin插件支持跨域请求
	g.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"authorization", "origin", "content-type", "accept", "timestamp", "sign", "token", "x-token"},
		ExposeHeaders: []string{"Content-Length", "Accept-Ranges", "Content-Range", "Content-Disposition"},
		// AllowCredentials: true,
	}))

	g.Use(mw...)

	r := g.Group("/api/v1")
	{
		mineRouter := r.Group("/mine")
		{
			mineRouter.POST("/login", handler.MineHandler.Login())
		}

		ruleRouter := r.Group("/rule")
		{
			ruleRouter.POST("/food/create", handler.RuleHandler.FoodCreate())
		}
	}

	return g
}
