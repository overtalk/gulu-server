package player

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/module"
)

type player struct {
	db module.DB
}

func Player(i *module.InternalService) {
	p := player{
		db: i.DB,
	}

	i.Gate.Add("POST", "/v1/user/pvp", p.SetPvpInfo, func(c *gin.Context) {
		fmt.Println("SetPvpInfo 后面的函数啊")
	})
	i.Gate.Add("POST", "/v1/user/info", p.SetBasicInfo)
	i.Gate.Add("POST", "/v1/user/common", p.CommonOp)
}
