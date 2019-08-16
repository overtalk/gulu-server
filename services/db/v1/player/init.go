package player

import (
	"database/sql"

	"github.com/didi/gendry/builder"

	"gitlab.com/SausageShoot/admin-server/model"
	"gitlab.com/SausageShoot/admin-server/module"
)

var (
	playerTable = "player"
)

type player struct {
	gm     module.GameGM
	db     *sql.DB
	pModel model.Player
	update map[string]interface{}
}

func Player(pl model.Player, db *sql.DB, gm module.GameGM) *player {
	return &player{
		gm:     gm,
		db:     db,
		pModel: pl,
		update: make(map[string]interface{}),
	}
}

// write data to db
func (p *player) Apply() error {
	where := map[string]interface{}{
		"id = ": p.pModel.Id,
	}

	cond, values, err := builder.BuildUpdate(playerTable, where, p.update)
	if err != nil {
		return err
	}
	_, err = p.db.Exec(cond, values...)
	return err
}
