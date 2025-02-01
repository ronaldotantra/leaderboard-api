package logger

import "errors"

// Fields ...
type Fields map[string]interface{}

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

// Logger ...
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
}

// A global variable so that log functions can be directly accessed
var log Logger

// SetRepository ...
func SetRepository(loggerInstance Logger) {
	log = loggerInstance
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}
