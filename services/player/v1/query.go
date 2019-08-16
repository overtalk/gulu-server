package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) Query(ctx context.Context, requestMessage interface{}) protocol.Response {
	return protocol.PostResponse{ErrCode: errtable.OkCode, ID: 200}
	req, ok := requestMessage.(*protocol.PlayerQuery)
	resp := protocol.PostResponse{ErrCode: errtable.OkCode}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return protocol.PostResponse{
			ErrCode: commonErrCode1,
			Msg:     "convert Query request error",
		}
	}

	if err := p.db.QueryPlayer(req); err != nil {
		resp.ErrCode = commonErrCode2
		resp.Msg = "Invalid PlayerID"
		return resp
	}

	resp.ID = req.ID

	return resp
}
