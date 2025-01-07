package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.Info("hello world")
	logrus.Error("hello world")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Warn("hello world")
	logrus.Error("hello world")
}
