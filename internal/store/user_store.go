package store

import (
	"context"
	"fmt"

	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/pkg/database"
	"gorm.io/gorm"
)

type UserStore interface {
	Save(ctx context.Context, u *model.User) error

	FindByEmail(ctx context.Context, email string) (*model.User, error)

	FindByUserID(ctx context.Context, userID int) (*model.User, error)
}

func NewUserStore(db *gorm.DB) (UserStore, error) {
	return &userStore{db: db}, nil
}

type userStore struct {
	db *gorm.DB
}

func (s *userStore) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var (
		db     = database.FromContext(ctx, s.db).WithContext(ctx)
		result model.User
	)
	if err := db.Where("email = ?", email).First(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, apierr.ErrResourceNotFound.WithMessage("user not found")
		}
		return nil, err
	}
	return &result, nil
}

func (s *userStore) FindByUserID(ctx context.Context, userID int) (*model.User, error) {
	var (
		db     = database.FromContext(ctx, s.db).WithContext(ctx)
		result model.User
	)

	if err := db.Where("user_id = ?", userID).First(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, apierr.ErrResourceNotFound.WithMessage("user not found")
		}
		return nil, err
	}

	return &result, nil
}

func (s *userStore) Save(ctx context.Context, u *model.User) error {
	if err := database.FromContext(ctx, s.db).
		WithContext(ctx).
		Save(u).Error; err != nil {
		fmt.Printf("failed")
		return apierr.ErrInternalServerError.WithMessage("user creation failed")
	}

	return nil
}
