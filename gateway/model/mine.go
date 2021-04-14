package model

import (
	"gengine/repository/model"
)

type (
	// 用户登录
	MineLoginRequest struct {
		HttpModel
		AccountNo string `json:"account_no"`
		Password  string `json:"password"`
	}
	UserLoginData struct {
		AccessToken  string       `json:"access_token"`
		RefreshToken string       `json:"refresh_token"`
		UserInfo     model.NfUser `json:"user_info"`
	}
	MineLoginResponse struct {
		BaseModel
		Data UserLoginData `json:"data"`
	}
)
