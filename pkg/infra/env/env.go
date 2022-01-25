package env

import (
	"github.com/gosidekick/goconfig"
	"log"
)

var (
	e *Environment
)

type Environment struct {
	AwsRegion   string `cfg:"AWS_REGION" cfgDefault:"us-east-1"`
	AwsSecret   string `cfg:"AWS_SECRET_ACCESS_KEY"`
	AwsKey      string `cfg:"AWS_ACCESS_KEY_ID"`
	AwsEndpoint string `cfg:"AWS_ENDPOINT"`

	DbHost    string `cfg:"DATABASE_HOST" cfgDefault:"localhost"`
	DbUser    string `cfg:"DATABASE_USER" cfgDefault:"user"`
	DbPass    string `cfg:"DATABASE_PASSWORD" cfgDefault:"passwd"`
	DbName    string `cfg:"DATABASE_NAME" cfgDefault:"schema"`
	DbConnStr string `cfg:"DATABASE_CONN_STR" cfgDefault:"%s:%s@tcp(%s)/%s?timeout=10s&tls=skip-verify"`
	DbPool    int    `cfg:"DATABASE_POOL" cfgDefault:"20"`

	HttpTimeout int    `cfg:"HTTP_TIMEOUT" cfgDefault:"10"`
	Port        string `cfg:"HTTP_PORT" cfgDefault:"8080"`

	Env      string `cfg:"ENVIRONMENT" cfgDefault:"local"`
	LogLevel string `cfg:"LOG_LEVEL" cfgDefault:"info"`
}

func Get() *Environment {
	if e == nil {
		e = &Environment{}
		err := goconfig.Parse(e)
		if err != nil {
			log.Fatal(err)
		}
	}
	return e
}

func SetEnvs(envs *Environment) {
	e = envs
}
