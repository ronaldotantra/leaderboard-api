package logrus

import (
	"os"

	"github.com/ronaldotantra/leaderboard-api/internal/logger"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{}
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

// NewLogrusLogger ...
func NewLogrusLogger(config *logger.Configuration) logger.Logger {
	level, err := logrus.ParseLevel(config.ConsoleLevel)
	if err != nil {
		panic(err)
	}
	ILogger := &logrus.Logger{
		Out:       os.Stdout,
		Level:     level,
		Formatter: getFormatter(config.ConsoleJSONFormat),
	}
	return &logrusLogger{
		logger: ILogger,
	}
}
