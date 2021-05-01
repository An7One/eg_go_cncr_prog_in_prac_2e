package base

import "github.com/an7one/eg_go_cncr_prog_in_prac_2e/src/util/log/field"

// Option option(s) of the logger
type Option interface {
	Name() string
}

// OptWithLocation a type of option of the logger
// with value of bool type
// representing whether the log includes where the code is called
type OptWithLocation struct {
	Value bool
}

func (opt OptWithLocation) Name() string {
	return "with location"
}

// MyLogger interface of the logger
type MyLogger interface {
	Name() string
	Level() LogLevel
	Format() LogFormat
	Options() []Option

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})

	// WithFields keeps track of surplus fields
	WithFields(fields ...field.Field) MyLogger
}
