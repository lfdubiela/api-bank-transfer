package logger

import (
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/env"
	"os"
	"time"

	logJSON "github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

var (
	log     *logJSON.Logger
	App     string
	Env     string
)

const (
	AppName = "api-bank-transfer"
	TimestampFormat = time.RFC3339

	FieldKeyTime = "logdate"
	FieldKeyMsg  = "messages"

	FieldApp = "app"
	FieldEnv = "env"
	FieldTid = "tid"
)

func init() {
	App = AppName
	Env = env.Get().Env

	log = logJSON.New()
	formatter := &logJSON.JSONFormatter{
		TimestampFormat: TimestampFormat,
		FieldMap: logJSON.FieldMap{
			logJSON.FieldKeyTime: FieldKeyTime,
			logJSON.FieldKeyMsg:  FieldKeyMsg,
		},
	}

	ll, err := logJSON.ParseLevel(env.Get().LogLevel)
	if err != nil {
		ll = logJSON.InfoLevel
	}

	log.SetLevel(ll)
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
}

func output(tid string, fields Fields) *logJSON.Entry {
	if fields == nil {
		fields = Fields{}
	}

	fields[FieldApp] = App
	fields[FieldEnv] = Env
	fields[FieldTid] = tid
	return log.WithFields(logJSON.Fields(fields))
}

func Debug(message string, cid string, fields Fields) {
	output(cid, fields).Debug(message)
}

func Info(message string, cid string, fields Fields) {
	output(cid, fields).Info(message)
}

func Warn(message string, cid string, fields Fields) {
	output(cid, fields).Warn(message)
}

func Fatal(message string, cid string, fields Fields) {
	output(cid, fields).Fatal(message)
}

func Error(message string, cid string, fields Fields) {
	output(cid, fields).Error(message)
}
