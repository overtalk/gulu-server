package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/didi/gendry/manager"

	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

// Pool is the constructor
func Connect(c Config) *sql.DB {
	conn, err := manager.New(c.DBName, c.Username, c.Password, c.Host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second),
	).Port(c.Port).Open(true)
	if err != nil {
		logger.Logger.Fatal(
			"Create MySQL Connection",
			logger.ErrorField(err),
			logger.Field("mySQL config", c))
	}
	return conn
}
