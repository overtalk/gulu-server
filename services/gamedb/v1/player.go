package gamedb

import (
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"

	"gitlab.com/SausageShoot/admin-server/model"
	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/services/gamedb/v1/player"
)

// get sql.GameDB
func (p *pool) GetPlayerByID(id int) (module.GamePlayer, error) {
	db, err := p.getDBByPlayerID(id)
	if err != nil {
		return nil, err
	}

	// query data
	where := map[string]interface{}{
		"id = ": id,
	}
	cond, values, err := builder.BuildSelect(playerTable, where, []string{"*"})
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(cond, values...)
	var pl model.Player
	if err := scanner.Scan(rows, &pl); err != nil {
		return nil, err
	}

	return player.Player(pl, db, p.gm), nil

}

func (p *pool) QueryPlayer(query *protocol.PlayerQuery) (*model.Player, error) {
	db, err := p.getDBByPlayerID(query.ID)
	if err != nil {
		return nil, err
	}

	// query data
	where := map[string]interface{}{
		"id = ": query.ID,
	}
	cond, values, err := builder.BuildSelect(playerTable, where, []string{"*"})
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(cond, values...)
	var pl model.Player
	if err := scanner.Scan(rows, &pl); err != nil {
		return nil, err
	}

	return &pl, err
}
