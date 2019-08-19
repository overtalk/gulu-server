package module

type InternalService struct {
	GameDB GameDB
	Gate   Gate
	GM     GameGM
	Cache  Cache
	DB     DB
}
