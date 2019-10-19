package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var selfLogger = logrus.New()

func SelfLogger() *logrus.Logger {
	if strings.ToLower(os.Getenv("LOGLEVEL")) != "error" {
		selfLogger.SetFormatter(&logrus.TextFormatter{})
		selfLogger.SetLevel(logrus.DebugLevel)
	} else {
		selfLogger.SetFormatter(&logrus.TextFormatter{})
		selfLogger.SetLevel(logrus.ErrorLevel)
	}

	selfLogger.SetReportCaller(true)
	return selfLogger
}
