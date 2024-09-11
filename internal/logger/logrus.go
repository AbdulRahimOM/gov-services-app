package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLoggerWithServiceName(serviceName string) *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	if os.Getenv("LOG_TO_FILE") == "true" {
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.SetLevel(logrus.InfoLevel)

	logEntry := log.WithField("service", serviceName)
	return logEntry
}
