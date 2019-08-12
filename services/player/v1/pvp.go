package player

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/common"
)

func (p *player) SetPvpInfo(c *gin.Context) {
	resp := common.Post{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	req := common.TurnPvpInfo(body)

	// todo: db operation
	pl, err := p.db.GetPlayerByID(req.ID)
	if err != nil {
		// todo: error handle
		resp.ErrCode = 1001
		resp.Msg = "没有找到player"
		c.String(200, resp.Encode())
		return
	}
	if err := pl.SetPvpInfo(req).Apply(); err != nil {

	}

	resp.ErrCode = 1000
	resp.Msg = "ok"

	c.String(200, resp.Encode())
	fmt.Println("all over")
}
