package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/controller"
	"github.com/merge/shopping-card/internal/handler"
	"github.com/merge/shopping-card/internal/handler/middleware"
	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/internal/store"
	"go.uber.org/fx"
)

type Server struct {
	server          *http.Server
	engine          *gin.Engine
	authController  *controller.AuthController
	appContoller    *controller.AppController
	adminController *controller.AdminController

	accessTokenStore store.AccessTokenStore
}

func NewServer(
	lc fx.Lifecycle,
	authController *controller.AuthController,
	appContoller *controller.AppController,
	adminController *controller.AdminController,
	accessTokenStore store.AccessTokenStore,

) (*Server, error) {

	srv := Server{
		engine:           gin.New(),
		authController:   authController,
		accessTokenStore: accessTokenStore,
		appContoller:     appContoller,
		adminController:  adminController,
	}

	srv.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: srv.engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("Start to rest api server :%d", 8000)
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

func (srv *Server) Router() error {

	v1 := srv.engine.Group("/api/v1")

	publicGroup := v1.Group("")

	publicGroup.POST("login", handler.Wrap(srv.authController.LoginHandler))
	publicGroup.POST("signup", handler.Wrap(srv.authController.SignupHandler))

	adminGroup := v1.Group("/admin", middleware.AuthMiddleware(srv.accessTokenStore,
		string(model.RoleAdmin)))

	adminGroup.POST("/add-item", handler.Wrap(srv.adminController.AddItemHandler))
	adminGroup.PUT("/disable-user", handler.Wrap(srv.adminController.DisableUserHandler))

	userGroup := v1.Group("/user", middleware.AuthMiddleware(srv.accessTokenStore,
		string(model.RoleUser)))

	userGroup.POST("/add-item", handler.Wrap(srv.appContoller.AddToCartHandler))
	userGroup.GET("/all-item", handler.Wrap(srv.appContoller.ListAllItemsHandler))
	userGroup.GET("/cart-item", handler.Wrap(srv.appContoller.ListCartHandler))
	userGroup.POST("/remove-item", handler.Wrap(srv.appContoller.RemoveItemFromCartHandler))

	return nil
}
