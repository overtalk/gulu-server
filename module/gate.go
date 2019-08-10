package module

import (
	"github.com/gin-gonic/gin"
)

type Gate interface {
	Start()
	Stop()
	Add(httpMethod, relativePath string, handlers ...gin.HandlerFunc)
}
