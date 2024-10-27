package utils

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func SetupLogger() {
	leg.SetFormatted(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}
