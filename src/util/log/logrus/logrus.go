package logrus

import (
	"an7one.com/eg_go_cncr_prog_in_prac_2e/src/util/log/base"

	"github.com/Sirupsen/logrus"
)

type loggerLogrus struct {
	level           base.LogLevel
	format          base.LogFormat
	optWithLocation base.OptWithLocation
	inner           *logrus.Entry
}

func NewLogger() base.MyLogger {
	return NewLoggerBy()
}

func NewLoggerBy() base.MyLogger {

}

func initInnerLogger() *logrus.Entry {

}
