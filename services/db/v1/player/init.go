package player

import (
	"database/sql"

	"github.com/didi/gendry/builder"

	"gitlab.com/SausageShoot/admin-server/model"
)

var (
	playerTable = "player"
)

type player struct {
	db     *sql.DB
	pModel model.Player
	update map[string]interface{}
}

func Player(pl model.Player, db *sql.DB) *player {
	return &player{
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
