package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) SetPvpInfo(ctx context.Context, requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.PvpInfo)
	resp := protocol.PostResponse{ErrCode: errtable.OkCode}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return protocol.PostResponse{
			ErrCode: pvpErrCode1,
			Msg:     "convert pvp request error",
		}
	}

	player, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		log.Logger.Error("get player",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: pvpErrCode2,
			Msg:     "get player error",
		}
	}

	if err := player.SetPvpInfo(req); err != nil {
		log.Logger.Error("set pvp info",
			log.ErrorField(err),
			log.Field("todo", req),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: pvpErrCode3,
			Msg:     "DAO ERROR(set pvp info)",
		}
	}

	if err := player.Apply(); err != nil {
		log.Logger.Error("player apply",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: pvpErrCode4,
			Msg:     "DAO ERROR(apply player after set pvp info)",
		}
	}

	return resp
}
