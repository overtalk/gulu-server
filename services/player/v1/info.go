package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) SetBasicInfo(ctx context.Context, requestMessage interface{}) interface{} {
	req, ok := requestMessage.(*protocol.BasicInfo)
	resp := protocol.PostResponse{Response: protocol.Response{ErrCode: errtable.OkCode}}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		resp.ErrCode = infoErrCode1
		resp.Msg = "convert basicInfo request error"
		return resp
	}

	player, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		log.Logger.Error("get player",
			log.ErrorField(err),
			log.Field("id", req.ID))
		resp.ErrCode = infoErrCode2
		resp.Msg = "get player error"
		return resp
	}

	player.SetBasicInfo(req)

	if err := player.Apply(); err != nil {
		log.Logger.Error("player apply",
			log.ErrorField(err),
			log.Field("id", req.ID))
		resp.ErrCode = infoErrCode3
		resp.Msg = "DAO ERROR(apply player after set basic info)"
	}

	return resp
}
