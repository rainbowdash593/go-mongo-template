package logging

import (
	log "github.com/sirupsen/logrus"
)

type Option func(*log.Logger)

func Level(lvl string) Option {
	return func(l *log.Logger) {
		logLevel, err := log.ParseLevel(lvl)
		if err != nil {
			logLevel = _logLevel
		}
		l.SetLevel(logLevel)
	}
}
