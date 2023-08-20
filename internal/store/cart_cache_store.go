package store

import (
	"context"
	"fmt"

	"github.com/merge/shopping-card/internal/handler/apierr"
	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/pkg/database"
	"gorm.io/gorm"
)

type CartCacheStore interface {
	Save(ctx context.Context, u *model.CartCache) error

	FindByUserID(ctx context.Context, ID string) ([]*model.CartCache, error)

	RemoveById(ctx context.Context, ID string, userID string) error
}

func NewCartCacheStore(db *gorm.DB) (CartCacheStore, error) {
	return &cartCacheStore{db: db}, nil
}

type cartCacheStore struct {
	db *gorm.DB
}

func (s *cartCacheStore) RemoveById(ctx context.Context, ID string, userID string) error {
	var (
		db = database.FromContext(ctx, s.db).WithContext(ctx)
	)
	if err := db.Delete(&model.CartCache{}, "cart_caches.user_id = ? and cart_caches.item_id = ?", userID, ID).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("failed")
		}
		return apierr.ErrInternalServerError.WithMessage("unable to process request at the moment")
	}

	return nil
}

// FindByID implements ItemStore.
func (s *cartCacheStore) FindByUserID(ctx context.Context, ID string) ([]*model.CartCache, error) {
	var (
		db     = database.FromContext(ctx, s.db).WithContext(ctx)
		result = []*model.CartCache{}
	)
	if err := db.Where("user_id = ?", ID).Find(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("failed")
		}
		return nil, err
	}
	return result, nil
}

// Save implements ItemStore.
func (s *cartCacheStore) Save(ctx context.Context, u *model.CartCache) error {
	if err := database.FromContext(ctx, s.db).
		WithContext(ctx).
		Save(u).Error; err != nil {
		fmt.Printf("failed")
	}
	return nil
}
