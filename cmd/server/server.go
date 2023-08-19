package main

import (
	"github.com/merge/shopping-card/internal/server"
	"go.uber.org/fx"
)

func runApplication() {
	fx.New(

		fx.Invoke(
			func(srv *server.Server) error {
				return srv.Router()
			}),
	).Run()
}
