package module

type Cache interface {
	// token
	GetPlayerIDByToken(token string) (int, error)
	SetToken(playerID int, token string) error
	UpdateExpire(token string) error
}
