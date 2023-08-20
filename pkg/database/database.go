package database

import (
	"context"
	"time"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type contextKey string

const dbKey = contextKey("db")

func FromContext(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if ctx == nil {
		return defaultDB
	}
	if db, ok := ctx.Value(dbKey).(*gorm.DB); ok {
		return db
	}
	return defaultDB
}

func WithContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

func Open() (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	for i := 0; i < 20; i++ {
		db, err = gorm.Open(gmysql.Open("root:password@tcp(127.0.0.1:3306)/merge?parseTime=true"), &gorm.Config{})
		if err == nil {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}

	DB, err := db.DB()
	if err != nil {
		return nil, err
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Minute * 30)

	return db, nil
}
