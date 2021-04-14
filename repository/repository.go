package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"

	"gengine/common"
)

type mysqlType uint8

const (
	MysqlType_Balance mysqlType = iota
	MysqlType_Master
)

const (
	_mgoDB = "zyx"
)

var (
	m           sync.RWMutex
	mysql       *gorm.DB
	mysqlMaster *gorm.DB
	mgoSession  *mgo.Session
)

var (
	UserRepo *userRepo

	MgoRuleRepo *mgoRuleRepo
)

func Init(cfg *common.Config) {
	logrus.Debug("repository Init 初始化仓储")

	// mysql
	UserRepo = NewUserRepo(cfg)

	// mongo
	MgoRuleRepo = NewMgoRuleRepo(cfg)
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

func GetMgoSession(cfg *common.Config) *mgo.Session {
	if mgoSession == nil {
		m.Lock()
		defer m.Unlock()
		if mgoSession == nil {
			mgoSession = connectMongo(cfg)
		}
	}
	return mgoSession
}

func connectMongo(cfg *common.Config) *mgo.Session {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:     cfg.Mongo.Hosts,
		Database:  "",
		Username:  cfg.Mongo.User,
		Password:  cfg.Mongo.Password,
		PoolLimit: 10480,
	})
	if err != nil {
		panic(err)
	}
	return s
}
