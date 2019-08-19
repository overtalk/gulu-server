package jwt

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const tokenKey = "Authorization"

var noAuthPath = []string{
	"/api/v1/login/account",
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	const useJWT = false
	if !useJWT {
		return func(c *gin.Context) {
			// 检测本路由是否需要进行token认证
			path := c.Request.URL.Path
			for _, p := range noAuthPath {
				if path == p {
					return
				}
			}
			fmt.Printf("path = %s,token = %s", path, c.GetHeader(tokenKey))
			// todo： check
		}
	}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
