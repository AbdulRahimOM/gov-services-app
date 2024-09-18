package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

// func NewLogger() *logrus.Logger {
// 	log := logrus.New()
// 	log.SetFormatter(&logrus.JSONFormatter{})
// 	log.SetOutput(os.Stdout)
// 	log.SetLevel(logrus.InfoLevel)
// 	return log
// }

func NewLoggerWithServiceName(serviceName string) *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	logEntry := log.WithField("service", serviceName)
	return logEntry
}
