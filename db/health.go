package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

func CheckHealth(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return sqlDB.PingContext(ctx)
}
