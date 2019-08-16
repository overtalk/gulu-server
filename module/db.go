package module

type User struct {
	Account  string
	Password string
	Auth     string
	Name     string
}

type DB interface {
	Users([]User, error)
}
