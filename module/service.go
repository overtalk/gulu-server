package module

type InternalService struct {
	DB   GameDB
	Gate Gate
	GM   GameGM
}
