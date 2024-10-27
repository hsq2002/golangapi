package utils

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func SetupLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}
