package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) SetBasicInfo(ctx context.Context, requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.BasicInfo)
	resp := protocol.PostResponse{ErrCode: errtable.OkCode}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return protocol.PostResponse{
			ErrCode: infoErrCode1,
			Msg:     "convert basicInfo request error",
		}
	}

	player, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		log.Logger.Error("get player",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: infoErrCode2,
			Msg:     "get player error",
		}
	}

	player.SetBasicInfo(req)

	if err := player.Apply(); err != nil {
		log.Logger.Error("player apply",
			log.ErrorField(err),
			log.Field("id", req.ID))
		return protocol.PostResponse{
			ErrCode: infoErrCode3,
			Msg:     "DAO ERROR(apply player after set basic info)",
		}
	}

	return resp
}
