package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/worker"
)

type AuthController struct {
	authWorker worker.AuthWorker
}

func NewAuthController(a worker.AuthWorker) (*AuthController, error) {
	c := AuthController{
		authWorker: a,
	}
	return &c, nil
}

func (a *AuthController) LoginHandler(gctx *gin.Context) (interface{}, error) {
	var (
		req LoginReq
	)

	if err := gctx.ShouldBind(&req); err != nil {
		return nil, apierr.ErrInvalidRequest.WithMessage(err.Error())
	}

	return a.authWorker.Login(gctx, req.Email, req.Password)

}

func (a *AuthController) SignupHandler(gctx *gin.Context) (interface{}, error) {
	var (
		req SignUpReq
	)

	if err := gctx.ShouldBind(&req); err != nil {
		return nil, apierr.ErrInvalidRequest.WithMessage(err.Error())
	}

	return a.authWorker.Signup(gctx, req.Email, req.Password)
	
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type SignUpReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}
