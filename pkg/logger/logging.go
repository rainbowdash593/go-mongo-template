package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	_logLevel = log.InfoLevel
)

var logger = log.Logger{}

func GetLogger() *log.Logger {
	return &logger
}

func ConfigureLogger(opts ...Option) *log.Logger {
	for _, opt := range opts {
		opt(&logger)
	}
	return &logger
}

func init() {
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(_logLevel)
	logger.SetReportCaller(true)
}
