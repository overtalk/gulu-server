package gate

import (
	"github.com/gin-gonic/contrib/static"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (g *gate) AddStatic(urlPrefix, root string, indexes bool) {
	if g.engine == nil {
		log.Logger.Fatal("add gin static file route",
			log.Field("urlPrefix", urlPrefix),
			log.Field("root", root),
			log.Field("indexes", indexes),
			log.ErrorField(nilEngineErr))
	}

	g.engine.Use(static.Serve(urlPrefix, static.LocalFile(root, indexes)))
}
