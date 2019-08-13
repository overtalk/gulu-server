package player

import (
	"fmt"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (p *player) CommonOP(requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.CommonOP)
	resp := protocol.PostResponse{ErrCode: 1000, Msg: "ok"}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return errtable.ConvertRequestErr
	}
	fmt.Println("CommonOP Request = ", req)

	return resp
}
