package db

import (
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"

	"gitlab.com/SausageShoot/admin-server/model"
	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/services/db/v1/player"
)

// get sql.DB
func (p *pool) GetPlayerByID(id int) (module.Player, error) {
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
