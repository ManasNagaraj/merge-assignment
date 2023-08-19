package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func NewServer(
	lc fx.Lifecycle,
) (*Server, error) {

	srv := Server{
		engine: gin.New(),
	}

	srv.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: srv.engine,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("starting a server brozz")
			return srv.Start()
		},
	})

	return &srv, nil
}

func (srv *Server) Start() error {
	go func() {
		err := srv.server.ListenAndServe()
		if err != nil {
			fmt.Println("server start failure")
		}
	}()
	return nil

}

func (s *Server) Router() error {

	return nil
}
