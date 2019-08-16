package gate

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

type gate struct {
	port   int
	srv    *http.Server
	engine *gin.Engine
}

// Gate is the constructor
func Gate(port int) *gate {
	e := gin.New()

	return &gate{
		port:   port,
		engine: e,
	}
}

// Start begin the gate server
func (g *gate) Start() {
	go func(g *gate) {
		g.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", g.port),
			Handler: g.engine,
		}

		log.Logger.Info("Server Start", log.Field("port", g.port))
		if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("Server Start", log.ErrorField(err))
		}
	}(g)
}

// Stop stop gate server gracefully
func (g *gate) Stop() {
	// prevent call stop without start
	if g.srv != nil {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := g.srv.Shutdown(ctx); err != nil {
			log.Logger.Fatal("Server Shutdown", log.ErrorField(err))
		}
	}
	log.Logger.Info("Server Exit")
}
