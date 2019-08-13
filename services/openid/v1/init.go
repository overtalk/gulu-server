package openid

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"

	"gitlab.com/SausageShoot/admin-server/model"
	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/utils/log"
	"gitlab.com/SausageShoot/admin-server/utils/mysql"
	"gitlab.com/SausageShoot/admin-server/utils/parse"
)

var nodesTable = "nodes"

func init() {
	scanner.SetTagName("json")
}

// OpenIdDB describes which node the openid binds to
type OpenIdDB struct {
	db        *sql.DB
	nodeConns map[int]*sql.DB
	address   string
}

func NewOpenIdDB(c mysql.Config) module.OpenIdDB {
	db := mysql.Connect(c)

	openIdConn := &OpenIdDB{
		db:        db,
		nodeConns: make(map[int]*sql.DB),
		address:   c.Host,
	}
	openIdConn.getAllNodes()

	return openIdConn
}

// getAllNodes get all nodes
func (conn *OpenIdDB) getAllNodes() {
	cond, values, err := builder.BuildSelect(nodesTable, nil, []string{"*"})
	if err != nil {
		log.Logger.Fatal("Get All Nodes", log.ErrorField(err))
	}

	rows, err := conn.db.Query(cond, values...)
	var nodes []model.Nodes
	if err := scanner.Scan(rows, &nodes); err != nil {
		log.Logger.Fatal("Scan All Nodes", log.ErrorField(err))
	}

	for _, v := range nodes {
		// split address to host and port to adjust frame `gendry`
		result := strings.Split(v.Address, ":")
		if len(result) != 2 {
			log.Logger.Fatal("Node Address",
				log.Field("error", fmt.Sprintf("invalid format : %s, expected[ip:port]", v.Address)))
		}

		conn.nodeConns[v.Id] = mysql.Connect(mysql.Config{
			Username: v.Username,
			Password: v.Password,
			Host:     result[0],
			DBName:   v.Database,
			Port:     int(parse.Int(result[1])),
			MaxConn:  0,
		})
	}
}
