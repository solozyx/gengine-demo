package service

import (
	"context"
	"gengine/common"
	. "gengine/gateway/model"
)

type IMineService interface {
	Login(req MineLoginRequest) (*MineLoginResponse, error)
}

func NewMineService(ctx context.Context) IMineService {
	userId, ok := ctx.Value(common.UserIdKey).(int)
	if !ok {
		return &mineService{}
	}
	return &mineService{
		UserId: userId,
	}
}

type mineService struct {
	UserId int
}

func (s *mineService) Login(req MineLoginRequest) (*MineLoginResponse, error) {
	//accountNo := strings.TrimSpace(req.AccountNo)
	//if accountNo == "" {
	//	return nil, errno.New(errno.ErrMustParaIsNull, errors.New("用户登录,必须填写手机号或邮箱 account_no"))
	//}
	//password := strings.TrimSpace(req.Password)
	//if password == "" {
	//	return nil, errno.New(errno.ErrMustParaIsNull, errors.New("用户登录,必须填写密码 password"))
	//}
	//logrus.Debugf("mineService Login req=%+v", req)
	//
	//hp := util.Sha1([]byte(password + common.UserPasswordSalt))
	//user, err := repository.NfUserRepo.SignIn(accountNo, hp)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// create token
	//accessToken, err := middleware.MakeToken(user.ID, 0, false)
	//if err != nil {
	//	logrus.Errorf("mineService Login make JWT accessToken error=%v", err.Error())
	//	return nil, errno.New(errno.ErrUserLoginFail, err)
	//}
	//refreshToken, err := middleware.MakeToken(user.ID, 0, true)
	//if err != nil {
	//	logrus.Errorf("mineService Login make JWT refreshToken error=%v", err.Error())
	//	return nil, errno.New(errno.ErrUserLoginFail, err)
	//}

	resp := &MineLoginResponse{}
	resp.Code = 0
	resp.Msg = "SUCCESS"

	resp.Data.AccessToken = req.AccountNo
	resp.Data.RefreshToken = req.Password
	//resp.Data.UserInfo = *user

	return resp, nil
}
