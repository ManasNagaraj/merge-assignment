package main

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/merge/shopping-card/internal/controller"
	"github.com/merge/shopping-card/internal/server"
	"github.com/merge/shopping-card/internal/store"
	"github.com/merge/shopping-card/internal/worker"
	"github.com/merge/shopping-card/pkg/database"
	"github.com/merge/shopping-card/pkg/logging"
	"go.uber.org/fx"
)

func runApplication() {
	fx.New(
		fx.Provide(
			//loggingInit
			logging.Get,

			//Database Init
			database.Open,
			store.NewUserStore,
			store.NewAccessTokenStore,
			store.NewCartCacheStore,
			store.NewItemStore,

			//worker Init
			worker.NewAuthWorker,
			worker.NewAdminWorker,
			worker.NewAppWorker,

			//Controllers Init
			controller.NewAuthController,
			controller.NewAdminController,
			controller.NewAppController,

			//Server Init
			server.NewServer,
		),
		fx.Invoke(
			func(srv *server.Server) error {
				return srv.Router()
			}),
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
