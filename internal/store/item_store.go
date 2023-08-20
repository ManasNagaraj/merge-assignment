package store

import (
	"context"
	"fmt"

	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/pkg/database"
	"gorm.io/gorm"
)

type ItemStore interface {
	Save(ctx context.Context, u *model.Item) error

	FindByID(ctx context.Context, ID string) (*model.Item, error)

	FindAll(ctx context.Context) ([]*model.Item, error)
}

func NewItemStore(db *gorm.DB) (ItemStore, error) {
	return &itemStore{db: db}, nil
}

type itemStore struct {
	db *gorm.DB
}

// FindAll implements ItemStore.
func (s *itemStore) FindAll(ctx context.Context) ([]*model.Item, error) {
	var (
		db     = database.FromContext(ctx, s.db).WithContext(ctx)
		result = []*model.Item{}
	)

	if err := db.Where("items.stock > 0").Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (s *itemStore) FindByID(ctx context.Context, ID string) (*model.Item, error) {
	var (
		db     = database.FromContext(ctx, s.db).WithContext(ctx)
		result model.Item
	)
	if err := db.Where("id = ?", ID).First(&result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("failed")
		}
		return nil, err
	}
	return &result, nil
}

func (s *itemStore) Save(ctx context.Context, u *model.Item) error {
	if err := database.FromContext(ctx, s.db).
		WithContext(ctx).
		Save(u).Error; err != nil {
		//return database.WrapError(err)
		fmt.Printf("failed")
	}
	return nil
}
