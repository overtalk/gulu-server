package errtable

import (
	"gitlab.com/SausageShoot/admin-server/protocol"
)

var (

	// ReadBodyErr defines get http body data from gin context error
	ReadBodyErr = protocol.PostResponse{
		ErrCode: 2001,
		Msg:     "read http body data error",
	}

	// ConvertRequestErr defines convert interface to service request error
	ConvertRequestErr = protocol.PostResponse{
		ErrCode: 2002,
		Msg:     "convert interface{} to service request error",
	}

	GetPlayerErr = protocol.PostResponse{
		ErrCode: 2003,
		Msg:     "get player error",
	}
)
