package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) CommonOP(ctx context.Context, requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.CommonOP)
	resp := protocol.PostResponse{ErrCode: errtable.OkCode}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return protocol.PostResponse{
			ErrCode: commonErrCode1,
			Msg:     "convert CommonOP request error",
		}
	}

	player, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		log.Logger.Error("get player",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: commonErrCode2,
			Msg:     "get player error",
		}
	}

	if err := player.CommonOP(req); err != nil {
		log.Logger.Error("set pvp info",
			log.ErrorField(err),
			log.Field("todo", req),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: commonErrCode3,
			Msg:     "DAO ERROR(set pvp info)",
		}
	}

	if err := player.Apply(); err != nil {
		log.Logger.Error("player apply",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: commonErrCode4,
			Msg:     "DAO ERROR(apply player after set basic info)",
		}
	}

	return resp
}
