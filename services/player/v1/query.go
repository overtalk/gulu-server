package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) Query(ctx context.Context, requestMessage interface{}) interface{} {
	req, ok := requestMessage.(*protocol.PlayerQuery)
	resp := protocol.PostResponse{Response: protocol.Response{ErrCode: errtable.OkCode}}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		resp.ErrCode = commonErrCode1
		resp.Msg = "convert Query request error"
		return resp
	}

	if err := p.db.QueryPlayer(req); err != nil {
		resp.ErrCode = commonErrCode2
		resp.Msg = "Invalid PlayerID"
		return resp
	}

	resp.ID = req.ID

	return resp
}
