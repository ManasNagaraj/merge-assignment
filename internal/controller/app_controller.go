package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/worker"
)

type AppController struct {
	appworker worker.AppWorker
}

func NewAppController(aw worker.AppWorker) (*AppController, error) {
	return &AppController{
		appworker: aw,
	}, nil
}

func (app *AppController) ListCartHandler(gctx *gin.Context) (interface{}, error) {
	return app.appworker.ListCart(gctx)
}

func (app *AppController) ListAllItemsHandler(gctx *gin.Context) (interface{}, error) {
	return app.appworker.ListAllItems(gctx)
}

func (app *AppController) AddToCartHandler(gctx *gin.Context) (interface{}, error) {

	var (
		req AddToCartReq
	)

	if err := gctx.ShouldBind(&req); err != nil {
		return nil, apierr.ErrInvalidRequest.WithMessage(err.Error())
	}
	return app.appworker.AddToCart(gctx, req.ItemID, req.Quantity)
}

func (app *AppController) RemoveItemFromCartHandler(gctx *gin.Context) (interface{}, error) {

	return app.appworker.RemoveItemFromCart(gctx, gctx.Param("id"), 1)
}

type AddToCartReq struct {
	ItemID   string `json:"item_id"`
	Quantity uint   `json:"quantity"`
}
