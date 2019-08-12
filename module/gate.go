package module

import (
	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/common"
)

type Handler func(req interface{}) common.Response

type Gate interface {
	Start()
	Stop()
	Add(httpMethod, relativePath string, handlers ...gin.HandlerFunc)
	Add1(httpMethod, relativePath string, handler Handler)
}
