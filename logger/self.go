package logger

import "github.com/sirupsen/logrus"

var selfLogger = logrus.New()

func SelfLogger() *logrus.Logger {
	if common.Mode() == "debug" {
		selfLogger.SetFormatter(&logrus.TextFormatter{})
		selfLogger.SetLevel(logrus.DebugLevel)
	} else {
		selfLogger.SetFormatter(&logrus.JSONFormatter{})
		selfLogger.SetLevel(logrus.ErrorLevel)
	}

	selfLogger.WithFields(logrus.Fields{"type": "handler"})
	selfLogger.SetReportCaller(true)
	return selfLogger
}
