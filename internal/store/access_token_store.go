package store

import (
	"context"
	"fmt"

	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/pkg/database"
	"gorm.io/gorm"
)

type AccessTokenStore interface {
	Save(ctx context.Context, u *model.AccessToken) error

	FindByUserID(ctx context.Context, ID string) (*model.AccessToken, error)

	FindByAccessToken(ctx context.Context, userId string) (*model.AccessToken, error)
}

func NewAccessTokenStore(db *gorm.DB) (AccessTokenStore, error) {
	return &accessTokenStore{db: db}, nil
}

type accessTokenStore struct {
	db *gorm.DB
}

func (s *accessTokenStore) FindByAccessToken(ctx context.Context, userId string) (*model.AccessToken, error) {
	var (
		db     = database.FromContext(ctx, s.db).WithContext(ctx)
		result model.AccessToken
	)
	if err := db.Where("access_token = ?", userId).First(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("failed")
		}
		return nil, err
	}
	return &result, nil
}

func (s *accessTokenStore) FindByUserID(ctx context.Context, userId string) (*model.AccessToken, error) {
	var (
		db     = database.FromContext(ctx, s.db).WithContext(ctx)
		result model.AccessToken
	)
	if err := db.Where("user_id = ?", userId).First(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("failed")
		}
		return nil, err
	}
	return &result, nil
}

func (s *accessTokenStore) Save(ctx context.Context, u *model.AccessToken) error {
	
	if err := database.FromContext(ctx, s.db).
		WithContext(ctx).
		Save(u).Error; err != nil {
		//return database.WrapError(err)
		fmt.Printf("failed")
	}
	return nil
}
