package worker

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/internal/store"
	"github.com/merge/shopping-card/pkg/utils/authutils"
)

type AuthWorker struct {
	userStore        store.UserStore
	accessTokenStore store.AccessTokenStore
}

func NewAuthWorker(u store.UserStore, ac store.AccessTokenStore) (AuthWorker, error) {
	a := AuthWorker{
		userStore:        u,
		accessTokenStore: ac,
	}
	return a, nil
}

func (a *AuthWorker) Login(ctx *gin.Context, email string, password string) (interface{}, error) {

	user, err := a.userStore.FindByEmail(ctx.Request.Context(), email)
	if err != nil {
		return nil, err
	}

	if err := authutils.MatchesPassword(user.Password, password); err != nil {
		return nil, apierr.ErrAuthenticationFail.WithMessage("Invalid Credentials")
	}

	if user.Disabled {
		return nil, apierr.ErrInvalidRequest.WithMessage("User Disabled")
	}

	refreshToken := authutils.GenerateRandomHex(64)
	accessToken := authutils.GenerateRandomHex(64)

	at := &model.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Active:       true,
		UserID:       user.UserID,
	}
	
	if err := a.accessTokenStore.Save(ctx, at); err != nil {
		return nil, apierr.ErrResourceConflict.WithMessagef("Failed to Sign In")
	}

	res := &SessionReponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	return res, nil
}

func (a *AuthWorker) Signup(ctx *gin.Context, email string, password string) (interface{}, error) {

	password, err := authutils.EncodePassword(password, 0)
	if err != nil {
		return nil, err
	}

	user := model.User{

		Email:    email,
		Password: password,
		Role:     string(model.RoleAdmin),
	}

	if err := a.userStore.Save(ctx, &user); err != nil {

		return nil, apierr.ErrResourceConflict.WithMessagef("Failed to create user")
	}

	refreshToken := authutils.GenerateRandomHex(64)
	accessToken := authutils.GenerateRandomHex(64)

	at := &model.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Role:         string(model.RoleAdmin),
		Active:       true,
		UserID:       user.UserID,
	}
	print("1 hello")
	if err := a.accessTokenStore.Save(ctx, at); err != nil {
		return nil, apierr.ErrResourceConflict.WithMessagef("Failed to Sign In")
	}

	res := &SessionReponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
	return res, nil
}

type SessionReponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
