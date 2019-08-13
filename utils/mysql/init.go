package mysql

import (
	"database/sql"
	"time"

	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"

	"gitlab.com/SausageShoot/admin-server/utils/log"
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
		log.Logger.Fatal(
			"Create MySQL Connection",
			log.ErrorField(err),
			log.Field("mySQL config", c))
	}
	return conn
}
