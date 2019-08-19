package player

import (
	"context"
	"fmt"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) SetPvpInfo(ctx context.Context, requestMessage interface{}) interface{} {
	req, ok := requestMessage.(*protocol.PvpInfo)
	resp := protocol.PostResponse{Response: protocol.Response{ErrCode: errtable.OkCode}}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		resp.ErrCode = pvpErrCode1
		resp.Msg = "convert pvp request error"
		return resp
	}

	fmt.Println(req)

	player, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		log.Logger.Error("get player",
			log.ErrorField(err),
			log.Field("id", req.ID))
		resp.ErrCode = pvpErrCode2
		resp.Msg = "get player error"
		return resp
	}

	if err := player.SetPvpInfo(req); err != nil {
		log.Logger.Error("set pvp info",
			log.ErrorField(err),
			log.Field("todo", req),
			log.Field("id", req.ID))
		resp.ErrCode = pvpErrCode3
		resp.Msg = "DAO ERROR(set pvp info)"
		return resp
	}

	if err := player.Apply(); err != nil {
		log.Logger.Error("player apply",
			log.ErrorField(err),
			log.Field("id", req.ID))
		resp.ErrCode = pvpErrCode4
		resp.Msg = "DAO ERROR(apply player after set pvp info)"
		return resp
	}

	return resp
}
