package common

import (
	"errors"
	"github.com/BurntSushi/toml"
	"path/filepath"
)

var (
	Cfg *Config
)

type Config struct {
	MySql struct {
		Balance struct {
			Host     string
			User     string
			Password string
			DataBase string
		}
		Master struct {
			Host     string
			User     string
			Password string
			DataBase string
		}
	}

	Mongo struct {
		Hosts    []string
		User     string
		Password string
	}

	Redis struct {
		Addrs    []string
		Password string
	}

	Mail struct {
		Username string
		Password string
		Host     string
		Port     string
	}

	Server struct {
		Host string
	}

	Nats struct {
		Addrs    []string
		User     string
		Password string
	}

	Cos struct {
		SecretId   string
		SecretKey  string
		BucketName string
		AppId      string
		Region     string
	}

	Obs struct {
		AccessKey  string
		SecretKey  string
		EndPoint   string
		BucketName string
	}

	Oss struct {
		AccessKeyID     string
		AccessKeySecret string
		EndPoint        string
		BucketName      string
	}
}

type Rds struct {
	Name         string
	User         string
	Password     string
	Server       string
	MaxIdleConns int `toml:"max_idle_conns"`
	MaxOpenConns int `toml:"max_open_conns"`
}

func InitConfig(path string, config interface{}) {
	if len(path) == 0 {
		panic(errors.New("configuration path is not provided"))
	}

	configPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	if _, err := toml.DecodeFile(configPath, config); err != nil {
		panic(err)
	}
}
