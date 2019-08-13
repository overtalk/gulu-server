package log_test

import (
	"fmt"
	"testing"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func TestLogger(t *testing.T) {
	err := fmt.Errorf("no password")
	log.Logger.Error("test",
		log.ErrorField(err),
		log.Field("key1", "value1"),
	)
}
