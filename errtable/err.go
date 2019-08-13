package errtable

import (
	"gitlab.com/SausageShoot/admin-server/protocol"
)

var (
	ResponseEncodeErr = `{"errcode":1003,"msg":"encode response error"}`

	// ReadBodyErr defines get http body data from gin context error
	ReadBodyErr = protocol.PostResponse{
		ErrCode: 1001,
		Msg:     "read http body data error",
	}

	// ConvertRequestErr defines convert interface to service request error
	ConvertRequestErr = protocol.PostResponse{
		ErrCode: 1002,
		Msg:     "convert interface{} to service request error",
	}
)
