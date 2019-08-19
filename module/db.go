package module

type Permission string

type DB interface {
	CheckPlayer(username, password string) (int, error)
	GetAuth(playerID int) Permission
}
