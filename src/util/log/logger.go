package log

import (
	"io"
	"sync"

	"an7one.com/eg_go_cncr_prog_in_prac_2e/src/util/log/base"
)

// LoggerCreator the creator of the logger
type LoggerCreator func(
	level base.LogLevel,
	format base.LogFormat,
	write io.Writer,
	options []base.Option) base.MyLogger

// loggerCreatorMap map of
var loggerCreatorMap = map[base.LoggerType]LoggerCreator{}

var rwm sync.RWMutex

func RegisterLogger() error {

}

func DLogger() base.MyLogger {

}

func Logger() base.MyLogger {

}
