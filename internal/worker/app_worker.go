package worker

import (
	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/internal/store"
)

type AppWorker struct {
	cartCacheStore store.CartCacheStore
	itemStore      store.ItemStore
}

func NewAppWorker(cc store.CartCacheStore, i store.ItemStore) (AppWorker, error) {
	a := AppWorker{
		cartCacheStore: cc,
		itemStore:      i,
	}
	return a, nil
}

func (app *AppWorker) ListAllItems(gctx *gin.Context) (interface{}, error) {
	result, err := app.itemStore.FindAll(gctx)
	if err != nil {
		return nil, apierr.ErrResourceNotFound.WithMessage("Not found")
	}

	res := &model.Response{
		Data:    result,
		Success: true,
	}
	return res, nil
}

func (app *AppWorker) ListCart(gctx *gin.Context) (interface{}, error) {
	userId := gctx.Value("userId")
	userIdStr, _ := userId.(string)
	result, err := app.cartCacheStore.FindByUserID(gctx, userIdStr)

	if err != nil {
		return nil, apierr.ErrInternalServerError.WithMessage("could not service right now")
	}

	res := &model.Response{
		Data:    result,
		Success: true,
	}

	return res, nil
}

func (app *AppWorker) AddToCart(gctx *gin.Context, itemId string, quantity uint) (interface{}, error) {
	userId := gctx.Value("userId")
	userIdStr, _ := userId.(string)
	cartItem := &model.CartCache{
		UserID:   string(userIdStr),
		ItemID:   itemId,
		Quantity: int(quantity),
	}

	if err := app.cartCacheStore.Save(gctx, cartItem); err != nil {
		return nil, apierr.ErrInternalServerError.WithMessage("unable to process request at the moment")
	}

	res := &model.Response{
		Success: true,
	}
	return res, nil

}

func (app *AppWorker) RemoveItemFromCart(gctx *gin.Context, itemId string, quantity uint) (interface{}, error) {
	userId := gctx.Value("userId")
	userIdStr, _ := userId.(string)

	if err := app.cartCacheStore.RemoveById(gctx,itemId,userIdStr); err != nil {
		return nil, apierr.ErrInternalServerError.WithMessage("unable to process request at the moment")
	}

	res := &model.Response{
		Success: true,
	}
	return res, nil

}
