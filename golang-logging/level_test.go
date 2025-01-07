package golang_logging

import (
	"github.com/sirupsen/logrus"
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
