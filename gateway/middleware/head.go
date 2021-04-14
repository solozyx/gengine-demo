package middleware

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gengine/common"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gengine/common/errno"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	NoSignUrls = []string{
		"/api/v1/static/",
		"/api/v1/video/upload",
		"/api/v1/app/file/upload",
		"/api/v1/app/file/info",
	}
)

func SignCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, noSign := range NoSignUrls {
			if strings.HasPrefix(ctx.GetString("request_uri"), noSign) {
				logrus.Debug("不验证sign")
				ctx.Next()
				return
			}
		}

		reqSign := strings.Trim(ctx.GetHeader("sign"), " ")
		if reqSign == "" {
			logrus.Warn("Header 参数校验失败.", " sign 不可为空. ")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errno.New(errno.ErrParas, errors.New("sign 不可为空")))
			return
		}

		reqMethod := ctx.Request.Method
		reqHost := ctx.Request.Host
		reqPath := ctx.Request.URL.Path
		reqTimestamp := ctx.GetHeader("timestamp")

		//rt, err := strconv.ParseInt(reqTimestamp, 10, 64)
		//if err != nil {
		//	logrus.Warnf("Header.timestamp 格式不正确. timestamp: %v", reqTimestamp)
		//	ctx.AbortWithStatusJSON(http.StatusBadRequest, errno.New(errno.ErrParas, errors.New("Header.timestamp 格式不正确")))
		//	return
		//}
		//if math.Abs(float64(rt-time.Now().Unix())) > 300 {
		//	logrus.Warnf("时间戳超过了5分钟.")
		//	ctx.AbortWithStatusJSON(http.StatusBadRequest, errno.New(errno.ErrParas, errors.New("Header.timestamp 超时")))
		//	return
		//}

		bodyStr := "{}"
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		if body != nil && len(body) > 0 {
			obj := make(map[string]interface{})
			if err := json.Unmarshal(body, &obj); err != nil {
				logrus.Errorf("json unmarshal request body fail. error: %v", err.Error())
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errno.New(errno.ErrParas, errors.New("请求体不是有效的json结构")))
				return
			}

			jsonStr, _ := json.Marshal(obj)
			bodyStr = string(jsonStr)
		}

		str := fmt.Sprintf("%s%s%s%s%s%s", reqMethod, reqHost, reqPath, bodyStr, reqTimestamp, common.HeaderSignSalt)
		logrus.Debug("sign before: ", str)
		h := md5.New()
		h.Write([]byte(str))
		bs := h.Sum(nil)
		secretStr := hex.EncodeToString(bs)

		logrus.Debugf("server sign: %s, \tuser sign: %s", secretStr, reqSign)

		// 验证
		if secretStr != reqSign {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errno.New(errno.ErrParas, errors.New("sign 校验未通过")))
			return
		}

		// 回写Body [读取一次Body后,Body会被清空,所以需要回写]
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		ctx.Next()
	}
}

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, timestamp, sign, token")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}
