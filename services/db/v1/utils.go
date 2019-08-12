package db

import (
	"database/sql"
)

// get sql.DB
func (p *pool) getDBByPlayerID(playerID int) (*sql.DB, error) {
	playerInfo, err := p.openidDB.GetPlayerInfoByPlayerID(playerID)
	if err != nil {
		return nil, err
	}
	return playerInfo.DB, nil
}
