package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hello Logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "agung")
	entry.Info("Hello Entry")
}
