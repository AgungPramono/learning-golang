package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Test Logger Formater")
	logger.Trace("Test Logger Trace")
	logger.Warn("Error")
}
