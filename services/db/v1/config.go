package db

// Config implement MySQL config
type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     int
}
