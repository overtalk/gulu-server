package logger_test

import (
	"fmt"
	"testing"

	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

func TestLogger(t *testing.T) {
	err := fmt.Errorf("no password")
	logger.Logger.Error("test",
		logger.ErrorField(err),
		logger.Field("key1", "value1"),
	)
}
