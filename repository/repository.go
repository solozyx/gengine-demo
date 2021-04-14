package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"

	"gengine/common"
)

type mysqlType uint8

const (
	MysqlType_Balance mysqlType = iota
	MysqlType_Master
)

var (
	m           sync.RWMutex
	mysql       *gorm.DB
	mysqlMaster *gorm.DB
)

var (
	NfUserRepo *userRepo
)

func Init(cfg *common.Config) {
	logrus.Debug("repository Init 初始化仓储")

	// mysql
	NfUserRepo = NewUserRepo(cfg)
}

func connectMysql(cfg *common.Config) *gorm.DB {
	dst := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySql.Balance.User, cfg.MySql.Balance.Password, cfg.MySql.Balance.Host, cfg.MySql.Balance.DataBase)
	db, err := gorm.Open("mysql", dst)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(128)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetConnMaxLifetime(time.Hour)

	// DEBUG
	db.LogMode(true)
	db.SetLogger(logrus.StandardLogger())

	return db
}

func connectMysqlMaster(cfg *common.Config) *gorm.DB {
	dst := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySql.Master.User, cfg.MySql.Master.Password, cfg.MySql.Master.Host, cfg.MySql.Master.DataBase)
	db, err := gorm.Open("mysql", dst)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(128)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetConnMaxLifetime(time.Hour)

	// DEBUG
	db.LogMode(true)
	db.SetLogger(logrus.StandardLogger())

	return db
}

func GetMysqlDB(cfg *common.Config, t mysqlType) *gorm.DB {
	switch t {
	case MysqlType_Balance:
		if mysql == nil || mysql.DB().Ping() != nil {
			m.Lock()
			defer m.Unlock()
			if mysql == nil || mysql.DB().Ping() != nil {
				mysql = connectMysql(cfg)
			}
		}
		return mysql
	case MysqlType_Master:
		if mysqlMaster == nil || mysqlMaster.DB().Ping() != nil {
			m.Lock()
			defer m.Unlock()
			if mysqlMaster == nil || mysqlMaster.DB().Ping() != nil {
				mysqlMaster = connectMysqlMaster(cfg)
			}
		}
		return mysqlMaster
	}
	return nil
}
