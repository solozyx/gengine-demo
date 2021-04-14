package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"gengine/common"
	"gengine/common/errno"
	"gengine/repository/model"
)

func NewUserRepo(cfg *common.Config) *userRepo {
	return &userRepo{
		db: GetMysqlDB(cfg, MysqlType_Balance),
	}
}

type userRepo struct {
	db *gorm.DB
}

func (p *userRepo) SignIn(accountNo, hp string) (*model.NfUser, error) {
	user := &model.NfUser{}
	err := p.db.Where("phone=? OR email=?", accountNo, accountNo).First(user).Error
	if err != nil {
		logrus.Errorf("userRepo SignIn accountNo=%s error=%v", accountNo, err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.ErrUserNotFound, errors.New(accountNo))
		}
		return nil, errno.New(errno.ErrDatabase, errors.New(accountNo))
	}
	logrus.Debugf("userRepo SignIn accountNo=%s user=%+v", accountNo, user)

	// 校验密码
	if user.Password != hp {
		logrus.Debugf("userRepo SignIn 密码校验未通过 数据库密码=%s 输入密码=%s", user.Password, hp)
		return user, errno.New(errno.ErrUserPassword, errors.New(accountNo))
	}

	logrus.Debugf("userRepo SignIn 密码校验通过 数据库密码=%s 输入密码=%s", user.Password, hp)
	return user, nil
}
