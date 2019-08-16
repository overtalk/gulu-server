package module

import (
	"database/sql"
)

type GameOpenIdDB interface {
	GetAllPlayerInfo() (map[string]*PlayerInfo, error)
	GetPlayerInfoByOpenId(openID string) (*PlayerInfo, error)
	GetPlayerInfoByPlayerID(PlayerID int) (*PlayerInfo, error)
	GetAllNode() map[int]*sql.DB
}

// PlayerInfo describes player info of the specified openid
type PlayerInfo struct {
	PlayerID int
	DB       *sql.DB
}
