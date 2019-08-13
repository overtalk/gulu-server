package player

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) Query(ctx context.Context, requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.PlayerQuery)
	resp := protocol.PostResponse{ErrCode: 1000, Msg: "ok"}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return errtable.ConvertRequestErr
	}

	req.ID = 1

	if err := p.db.QueryPlayer(req); err != nil {
		resp.ErrCode = 1001
		resp.Msg = "Invalid PlayerID"
		return resp
	}

	resp.ID = req.ID

	return resp
}
