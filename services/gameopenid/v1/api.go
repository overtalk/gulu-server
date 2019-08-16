package gameopenid

import (
	"database/sql"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"

	"gitlab.com/SausageShoot/admin-server/model"
	"gitlab.com/SausageShoot/admin-server/module"
)

var openidMappingTable = "openid_mapping"

func (conn *OpenIdDB) GetAllPlayerInfo() (map[string]*module.PlayerInfo, error) {
	cond, values, err := builder.BuildSelect(openidMappingTable, nil, []string{"*"})
	if err != nil {
		return nil, err
	}

	rows, err := conn.db.Query(cond, values...)
	var mapping []model.Openid_mapping
	if err := scanner.Scan(rows, &mapping); err != nil {
		return nil, err
	}

	res := make(map[string]*module.PlayerInfo)
	for _, v := range mapping {
		res[v.OpenId] = &module.PlayerInfo{
			PlayerID: v.PlayerId,
			DB:       conn.nodeConns[v.NodeId],
		}
	}
	return res, nil
}

func (conn *OpenIdDB) GetPlayerInfoByOpenId(openID string) (*module.PlayerInfo, error) {
	where := map[string]interface{}{
		"open_id = ": openID,
	}
	cond, values, err := builder.BuildSelect(openidMappingTable, where, []string{"*"})
	if err != nil {
		return nil, err
	}

	rows, err := conn.db.Query(cond, values...)
	var mapping model.Openid_mapping
	if err := scanner.Scan(rows, &mapping); err != nil {
		return nil, err
	}

	return &module.PlayerInfo{
		PlayerID: mapping.PlayerId,
		DB:       conn.nodeConns[mapping.NodeId],
	}, nil
}

func (conn *OpenIdDB) GetPlayerInfoByPlayerID(strPlayerID int) (*module.PlayerInfo, error) {
	where := map[string]interface{}{
		"player_id = ": strPlayerID,
	}
	cond, values, err := builder.BuildSelect(openidMappingTable, where, []string{"*"})
	if err != nil {
		return nil, err
	}

	rows, err := conn.db.Query(cond, values...)
	var mapping model.Openid_mapping
	if err := scanner.Scan(rows, &mapping); err != nil {
		return nil, err
	}

	return &module.PlayerInfo{
		PlayerID: mapping.PlayerId,
		DB:       conn.nodeConns[mapping.NodeId],
	}, nil
}

func (conn *OpenIdDB) GetAllNode() map[int]*sql.DB {
	return conn.nodeConns
}
