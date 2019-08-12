package player

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/req"
)

func (p *player) SetPvpInfo(c *gin.Context) {
	resp := req.PostResponse{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	req := req.PvpInfoReq(body)

	fmt.Println("PvpInfoReq = ", req)

	// todo: db operation

	resp.ErrCode = 1000
	resp.Msg = "ok"

	c.String(200, resp.Encode())
}