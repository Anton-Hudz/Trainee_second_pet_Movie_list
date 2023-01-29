package logger

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

func CustomLogger(loglevel string) error {
	var loggerLevel logrus.Level
	switch strings.ToLower(loglevel) {
	case "fatal":
		loggerLevel = logrus.FatalLevel
	case "error":
		loggerLevel = logrus.ErrorLevel
	case "warn":
		loggerLevel = logrus.WarnLevel
	case "info":
		loggerLevel = logrus.InfoLevel
	case "debug":
		loggerLevel = logrus.DebugLevel
	default:
		return errors.New("error of log level\n")
	}

	logrus.SetLevel(loggerLevel)
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	return nil
}
