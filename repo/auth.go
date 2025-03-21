package repo

import (
	"context"
	"errors"
	"time"

	"git.sindadsec.ir/asm/backend/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, org *models.Organization, rctx context.Context) error {
	var mysqlErr *mysql.MySQLError

	ctx, cancel := context.WithTimeout(rctx, time.Second*3)
	defer cancel()

	tx := db.Begin()
	if err := tx.WithContext(ctx).Create(org).Error; err != nil {
		tx.Rollback()
		errors.As(err, &mysqlErr)
		switch mysqlErr.Number {
		case 1062:
			return ErrDuplicateEntry
		default:
			return err
		}
	}

	tx.Commit()
	return nil
}
