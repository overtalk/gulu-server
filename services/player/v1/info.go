package player

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/common"
)

func (p *player) SetBasicInfo(c *gin.Context) {
	resp := common.Post{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	req := common.TurnBasicInfo(body)

	fmt.Println("BasicInfoReq = ", req)

	// todo: set db

	resp.ErrCode = 1000
	resp.Msg = "ok"

	c.String(200, resp.Encode())
}
