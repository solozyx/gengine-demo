/*
 *  GENERATE BY CODE. CAN NOT EDIT!!!
 */

package handler

import (
	"gengine/common/errno"
	. "gengine/gateway/model"
	"gengine/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ruleHandler struct{}

func (p *ruleHandler) FoodCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqModel RuleFoodCreateRequest
		if err := ctx.Bind(&reqModel); err != nil {
			log.Error(err)
			v := errno.New(errno.ErrBind, err)
			ctx.JSON(http.StatusBadRequest, BaseModel{
				Code: v.Code,
				Msg:  v.Error(),
			})
			return
		}
		s := service.NewRuleService(ctx)

		reqModel.Path = ctx.Request.RequestURI
		reqModel.Method = ctx.Request.Method
		reqModel.IP = ctx.ClientIP()

		resp, err := s.FoodCreate(reqModel)
		if err != nil {
			log.Error(err)
			if v, ok := err.(*errno.Err); ok {
				ctx.JSON(http.StatusBadRequest, BaseModel{
					Code: v.Code,
					Msg:  v.Error(),
				})
				return
			}

			v := errno.New(errno.InternalServerError, err)
			ctx.JSON(http.StatusInternalServerError, BaseModel{
				Code: v.Code,
				Msg:  v.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, resp)
		return
	}
}

func (p *ruleHandler) FoodCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqModel RuleFoodCheckRequest
		if err := ctx.Bind(&reqModel); err != nil {
			log.Error(err)
			v := errno.New(errno.ErrBind, err)
			ctx.JSON(http.StatusBadRequest, BaseModel{
				Code: v.Code,
				Msg:  v.Error(),
			})
			return
		}
		s := service.NewRuleService(ctx)

		reqModel.Path = ctx.Request.RequestURI
		reqModel.Method = ctx.Request.Method
		reqModel.IP = ctx.ClientIP()

		resp, err := s.FoodCheck(reqModel)
		if err != nil {
			log.Error(err)
			if v, ok := err.(*errno.Err); ok {
				ctx.JSON(http.StatusBadRequest, BaseModel{
					Code: v.Code,
					Msg:  v.Error(),
				})
				return
			}

			v := errno.New(errno.InternalServerError, err)
			ctx.JSON(http.StatusInternalServerError, BaseModel{
				Code: v.Code,
				Msg:  v.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, resp)
		return
	}
}
