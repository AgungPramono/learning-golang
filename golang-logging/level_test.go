package golang_logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("Trace level")
	logger.Debug("Debug Level")
	logger.Info("Info Level")
	logger.Warn("Warn Level")
	logger.Error("Error Level")
	logger.Fatal("Fatal Level")
}

func TestLoggingLevel(testing *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)

	logger.Trace("Trace level")
	logger.Debug("Debug Level")
	logger.Info("Info Level")
	logger.Warn("Warn Level")
	logger.Error("Error Level")
	logger.Fatal("Fatal Level")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Test Log Ke file")
	logger.Warn("warning")
	logger.Error("Error")
}
