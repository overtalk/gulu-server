package mysql

// Config implement MySQL config
type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     int
	MaxConn  int //最大连接数
}
