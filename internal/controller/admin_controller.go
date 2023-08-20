package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/worker"
)

type AdminController struct {
	adminWorker worker.AdminWorker
}

func NewAdminController(a worker.AdminWorker) (*AdminController, error) {
	c := AdminController{
		adminWorker: a,
	}

	return &c, nil
}

func (a *AdminController) DisableUserHandler(gctx *gin.Context) (interface{}, error) {
	var (
		req DisableUserReq
	)

	if err := gctx.ShouldBind(&req); err != nil {
		return nil, apierr.ErrInvalidRequest.WithMessage(err.Error())
	}

	return a.adminWorker.DisableUser(gctx, req.UserID, req.Message)
}

func (a *AdminController) AddItemHandler(gctx *gin.Context) (interface{}, error) {
	var (
		req AddItemReq
	)

	if err := gctx.ShouldBind(&req); err != nil {
		return nil, apierr.ErrInvalidRequest.WithMessage(err.Error())
	}

	return a.adminWorker.AddItem(gctx, req.Desc, req.Name, req.Quantity, req.Price)
}

type DisableUserReq struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type AddItemReq struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
}
