package player

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/module"
)

func (p *player) CommonOp(c *gin.Context) {
	resp := module.PostResponse{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	req := module.CommonOpReq(body)

	fmt.Println("CommonOpReq = ", req)

	// todo: db operation

	resp.ErrCode = 1000
	resp.Msg = "ok"

	c.String(200, resp.Encode())
}
