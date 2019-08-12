package openid_test

import (
	"testing"
)

func TestGetAllNode(t *testing.T) {
	openidDB := getConn(t)
	nodes := openidDB.GetAllNode()
	t.Log(nodes)
}

func TestGetAllPlayerInfo(t *testing.T) {
	openidDB := getConn(t)
	playerInfos, err := openidDB.GetAllPlayerInfo()
	if err != nil {
		t.Error(err)
		return
	}
	for k, v := range playerInfos {
		t.Log(k, v)
	}
}

func TestGetPlayerInfoByOpenId(t *testing.T) {
	openid := "qinhan"
	openidDB := getConn(t)
	playerInfo, err := openidDB.GetPlayerInfoByOpenId(openid)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(playerInfo)
}

func TestGetPlayerInfoByPlayerID(t *testing.T) {
	playerID := 1
	openidDB := getConn(t)
	playerInfo, err := openidDB.GetPlayerInfoByPlayerID(playerID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(playerInfo)
}
