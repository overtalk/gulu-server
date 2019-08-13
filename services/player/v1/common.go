package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) CommonOP(ctx context.Context, requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.CommonOP)
	resp := protocol.PostResponse{ErrCode: 1000, Msg: "ok"}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return errtable.ConvertRequestErr
	}

	player, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		log.Logger.Error("get player",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return errtable.GetPlayerErr
	}

	if err := player.CommonOP(req); err != nil {
		log.Logger.Error("set pvp info",
			log.ErrorField(err),
			log.Field("todo", req),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: 1001,
			Msg:     "DAO ERROR(set pvp info)",
		}
	}

	if err := player.Apply(); err != nil {
		log.Logger.Error("player apply",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: 1001,
			Msg:     "DAO ERROR(apply player after set basic info)",
		}
	}

	return resp
}
