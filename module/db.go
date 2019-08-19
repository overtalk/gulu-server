package module

type User struct {
	ID       int64  `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
}

type DB interface {
	CheckUser(username, password string) (int64, error)
	GetUser(id int64) (*User, error)
}
