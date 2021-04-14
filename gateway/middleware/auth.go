package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"gengine/common"
	"gengine/common/errno"
)

var (
	NoAuthUrls = []string{
		"/api/v1/static/",
		"/api/v1/mine/exist",
		"/api/v1/mine/emailcode",
		"/api/v1/mine/confirmcode",
		"/api/v1/mine/changepassword",
		"/api/v1/mine/login",
		"/api/v1/video/video",
		"/api/v1/video/getkeywords",
		"/api/v1/app/file/upload",
		"/api/v1/app/file/info",
	}
)

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, noAuth := range NoAuthUrls {
			if strings.HasPrefix(ctx.GetString("request_uri"), noAuth) {
				logrus.Debugf("request uri=%s JWT 不验证token", ctx.GetString("request_uri"))
				if strings.TrimSpace(ctx.GetHeader("token")) != "" {
					token := ctx.GetHeader("token")
					logrus.Debugf("request token=%s", token)
					pl, err := ValidateToken(token, false)
					if err == nil {
						ctx.Set(common.UserIdKey, pl.UserID)
					}
				}
				ctx.Next()
				return
			}
		}

		logrus.Debugf("request uri=%s JWT 验证token", ctx.GetString("request_uri"))

		token := ctx.GetHeader("token")
		logrus.Debugf("request token=%s", token)

		pl, err := ValidateToken(token, false)
		if err != nil {
			if err == jwt.ErrExpValidation {
				logrus.Warnf("token is expire. error=%v", err.Error())
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errno.ErrTokenExpire)
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errno.ErrToken)
			return
		}

		ctx.Set(common.UserIdKey, pl.UserID)
		ctx.Next()
	}
}

type CustomPayload struct {
	jwt.Payload
	UserID int `json:"user_id,omitempty"`
	Level  int `json:"level,omitempty"`
}

var tokenSecret = jwt.NewHS256([]byte(common.JwtTokenSecret))

const (
	ACCESSTOKENEXPIRY  = 24 * time.Hour
	REFRESHTOKENEXPIRY = 30 * 24 * time.Hour
)

func MakeToken(userID int, level int, isRefresh bool) (string, error) {
	now := time.Now()
	var pl *CustomPayload
	if isRefresh {
		pl = &CustomPayload{
			Payload: jwt.Payload{
				Issuer:         common.JwtPayloadIssuer,
				Subject:        "RefreshToken",
				Audience:       jwt.Audience{common.JwtPayloadAudience},
				ExpirationTime: jwt.NumericDate(now.Add(REFRESHTOKENEXPIRY)),
				NotBefore:      jwt.NumericDate(now.Add(-5 * time.Minute)),
				IssuedAt:       jwt.NumericDate(now),
				JWTID:          uuid.New().String(),
			},
			UserID: userID,
			Level:  level,
		}
	} else {
		pl = &CustomPayload{
			Payload: jwt.Payload{
				Issuer:         "secondchase.gateway",
				Subject:        "AccessToken",
				Audience:       jwt.Audience{common.JwtPayloadAudience},
				ExpirationTime: jwt.NumericDate(now.Add(ACCESSTOKENEXPIRY)),
				NotBefore:      jwt.NumericDate(now.Add(-5 * time.Minute)),
				IssuedAt:       jwt.NumericDate(now),
				JWTID:          uuid.New().String(),
			},
			UserID: userID,
			Level:  level,
		}
	}

	tokenBytes, err := jwt.Sign(pl, tokenSecret, jwt.ContentType("JWT"))
	if err != nil {
		logrus.Errorf("create JWT token fail. error: %v", err.Error())
		return "", err
	}

	return string(tokenBytes), nil
}

func ValidateToken(token string, isRefresh bool) (*CustomPayload, error) {
	now := time.Now()
	var pl CustomPayload

	expValidator := jwt.ExpirationTimeValidator(now)
	var subValidator jwt.Validator
	if isRefresh {
		subValidator = jwt.SubjectValidator("RefreshToken")
	} else {
		subValidator = jwt.SubjectValidator("AccessToken")
	}
	nbfValidator := jwt.NotBeforeValidator(now)
	audValidator := jwt.AudienceValidator(jwt.Audience{common.JwtPayloadAudience})
	validatePayload := jwt.ValidatePayload(&pl.Payload, expValidator, subValidator, nbfValidator, audValidator)

	_, err := jwt.Verify([]byte(token), tokenSecret, &pl, validatePayload)
	if err != nil {
		logrus.Errorf("parse JWT token fail. error: %v", err.Error())
		return &pl, err
	}

	return &pl, nil
}
