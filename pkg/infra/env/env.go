package env

import (
	"github.com/gosidekick/goconfig"
	"log"
)

var (
	e *Environment
)

type Environment struct {
	AwsRegion       string `cfg:"AWS_REGION" cfgDefault:"us-east-1"`
	AwsSecret       string `cfg:"AWS_SECRET_ACCESS_KEY"`
	AwsKey          string `cfg:"AWS_ACCESS_KEY_ID"`
	AwsEndpoint     string `cfg:"AWS_ENDPOINT"`
	AwsWebTokenFile string `cfg:"AWS_WEB_IDENTITY_TOKEN_FILE"`

	DbReadHost  string `cfg:"DATABASE_READ_HOST" cfgDefault:"localhost"`
	DbReadUser  string `cfg:"DATABASE_READ_USER" cfgDefault:"user"`
	DbReadPass  string `cfg:"DATABASE_READ_PASSWORD" cfgDefault:"passwd"`
	DbReadName  string `cfg:"DATABASE_READ_NAME" cfgDefault:"schema"`
	DbReadPool  int    `cfg:"DATABASE_READ_POOL" cfgDefault:"20"`
	DbWriteHost string `cfg:"DATABASE_WRITE_HOST" cfgDefault:"localhost"`
	DbWriteUser string `cfg:"DATABASE_WRITE_USER" cfgDefault:"user"`
	DbWritePass string `cfg:"DATABASE_WRITE_PASSWORD" cfgDefault:"passwd"`
	DbWriteName string `cfg:"DATABASE_WRITE_NAME" cfgDefault:"schema"`
	DbWritePool int    `cfg:"DATABASE_WRITE_POOL" cfgDefault:"20"`
	DbConnStr   string `cfg:"DATABASE_CONN_STR" cfgDefault:"%s:%s@tcp(%s)/%s?timeout=10s&tls=skip-verify"`

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
