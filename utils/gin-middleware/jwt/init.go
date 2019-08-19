package jwt

import (
	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/protocol"
)

const (
	TOKENKEY = "Authorization"
	ID       = "ID"
)

var noAuthPath = []string{
	"/api/v1/login/account",
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检测本路由是否需要进行token认证
		path := c.Request.URL.Path
		for _, p := range noAuthPath {
			if path == p {
				return
			}
		}

		// 解析token信息
		info, err := NewJWT().ParseToken(c.GetHeader(TOKENKEY))
		if err != nil {
			c.JSON(200, protocol.Response{ErrCode: 2222222, Msg: "token错"})
			c.Abort()
			return
		}

		// 设置playerID
		c.Set("ID", info.ID)
	}
}
