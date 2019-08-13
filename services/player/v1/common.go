package player

import (
	"fmt"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

func (p *player) CommonOP(requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.CommonOP)
	resp := protocol.PostResponse{ErrCode: 1000, Msg: "ok"}

	if !ok {
		logger.Logger.Error("Convert", logger.Field("request", requestMessage))
		return errtable.ConvertRequestErr
	}
	fmt.Println("CommonOP Request = ", req)

	return resp
}
