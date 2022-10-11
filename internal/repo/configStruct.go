package repo

import (
	"github.com/BurntSushi/toml"
)

var conf *Config

type Config struct {
	Host           string
	Postgresqlport int
	User           string
	Password       string
	DBname         string
	Port           string
}

func GetConfig() (*Config, error) {
	var err error
	if conf == nil {
		_, err = toml.DecodeFile("../configs/config.toml", &conf)
	}
	return conf, err
}

func GetPort() string {
	return conf.Port
}
