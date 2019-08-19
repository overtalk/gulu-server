package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) CommonOP(ctx context.Context, requestMessage interface{}) interface{} {
	req, ok := requestMessage.(*protocol.CommonOP)
	resp := protocol.PostResponse{Response: protocol.Response{ErrCode: errtable.OkCode}}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		resp.ErrCode = commonErrCode1
		resp.Msg = "convert CommonOP request error"
		return resp
	}

	player, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		log.Logger.Error("get player",
			log.ErrorField(err),
			log.Field("id", req.ID))
		resp.ErrCode = commonErrCode2
		resp.Msg = "get player error"
		return resp
	}

	switch req.OP {
	case 1:
		err = player.GetAllWeapons(req.ID)
	case 2:
		err = player.ClearWeapons(req.ID)
	case 3:
		err = player.GetAllFashions(req.ID)
	case 4:
		err = player.ClearFashions(req.ID)
	case 5:
		err = player.PveUnlock(req.ID)
	}

	if err != nil {
		log.Logger.Error("common op",
			log.ErrorField(err),
			log.Field("todo", req),
			log.Field("id", req.ID))
		resp.ErrCode = commonErrCode3
		resp.Msg = "DAO ERROR(common op)"
		return resp
	}

	if err := player.Apply(); err != nil {
		log.Logger.Error("player apply",
			log.ErrorField(err),
			log.Field("id", req.ID))
		resp.ErrCode = commonErrCode4
		resp.Msg = "DAO ERROR(apply player after common op)"
		return resp
	}

	return resp
}
