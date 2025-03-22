package repo

import (
	"context"
	"errors"
	"time"

	"git.sindadsec.ir/asm/backend/models"
	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string, rctx context.Context) (*models.User, error) {
	var user models.User

	ctx, cancel := context.WithTimeout(rctx, time.Second*5)
	defer cancel()

	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
