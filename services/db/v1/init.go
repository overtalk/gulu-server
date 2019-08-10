package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/didi/gendry/manager"
	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

type pool struct {
	conn *sql.DB
}

// Pool is the constructor
func Pool(c Config) *pool {
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
	return &pool{conn: conn}
}
