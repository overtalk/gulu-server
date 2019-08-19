package errtable

// 错误模块
type ErrModule int

const (
	Gate     ErrModule = 0
	DB       ErrModule = 1
	Auth     ErrModule = 2
	GM       ErrModule = 3
	OpenIdDB ErrModule = 4
	Player   ErrModule = 5

	GateString     string = "Gate"
	DBString       string = "GameDB"
	AuthString     string = "Auth"
	GMString       string = "GM"
	OpenIdDBString string = "OpenIdDB"
	PlayerString   string = "Player"
)
