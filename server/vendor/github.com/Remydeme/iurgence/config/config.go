package config

import (
	"errors"
	"github.com/jinzhu/configor"
	"os"
)

var Main Configuration

var (
	NO_ENV_VAR_FOUNDED = errors.New("No env variable founded for this key")
)

type Database struct {
	Dialect  string `default:"MangoDB"`
	Debug    bool   `default:"false"`
	Username string `required:"true" default:"root"`
	Password string `required:"true"`
	Host     string `required:"true"`
	Port     int
	DBname   string `required:"true"`
	SSLMode  bool   `default:"true"`
}

type JWT struct {
	Secret string `required:"true"`
}

type Server struct {
	Port   string `required:"true" default:":8080"`
	Domain string `required:"false"`
	Host   string `required:"false"`
}

type Configuration struct {
	Database Database `required:"true"`
	JWT      JWT      `required:"true"`
	Server   Server   `required:"true"`
}

func init() {

	PATH_CONF_FILE := os.Getenv("CONFIG_PATH")

	if PATH_CONF_FILE == "" {
		panic(NO_ENV_VAR_FOUNDED.Error())
	}

	if err := configor.Load(&Main, PATH_CONF_FILE); err != nil {
		panic(err.Error())
	}
}
