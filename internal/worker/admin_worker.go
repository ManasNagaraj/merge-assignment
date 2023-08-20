package worker

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/internal/store"
)

type AdminWorker struct {
	userStore store.UserStore
	itemStore store.ItemStore
}

func NewAdminWorker(u store.UserStore, i store.ItemStore) (AdminWorker, error) {
	a := AdminWorker{
		userStore: u,
		itemStore: i,
	}
	return a, nil
}

func (w *AdminWorker) DisableUser(gctx *gin.Context, userID int, message string) (interface{}, error) {
	user, err := w.userStore.FindByUserID(gctx, userID)
	if err != nil {
		return nil, err
	}

	user.Disabled = true

	if err := w.userStore.Save(gctx, user); err != nil {

		return nil, apierr.ErrResourceConflict.WithMessagef("Failed to create user")
	}

	return nil, nil
}

func (w *AdminWorker) AddItem(gctx *gin.Context, name string, desc string, quantity uint, price uint) (interface{}, error) {
	i := &model.Item{
		Name:      name,
		Desc:      desc,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Stock:     quantity,
		Price:     price,
		Disabled:  false,
	}
	if err := w.itemStore.Save(gctx, i); err != nil {
		return nil, apierr.ErrInvalidRequest.WithMessage("Unable to process your request")
	}

	return nil, nil
}
