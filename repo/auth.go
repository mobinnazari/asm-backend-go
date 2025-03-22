package repo

import (
	"context"
	"errors"
	"time"

	"git.sindadsec.ir/asm/backend/hash"
	"git.sindadsec.ir/asm/backend/mail"
	"git.sindadsec.ir/asm/backend/models"
	"github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, client *redis.Client, org *models.Organization, user *models.User, rctx context.Context, dialer *gomail.Dialer) error {
	var mysqlErr *mysql.MySQLError

	ctx, cancel := context.WithTimeout(rctx, time.Second*5)
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

	user.OrganizationID = org.ID
	if err := tx.WithContext(ctx).Create(user).Error; err != nil {
		tx.Rollback()
		errors.As(err, &mysqlErr)
		switch mysqlErr.Number {
		case 1062:
			return ErrDuplicateEntry
		default:
			return err
		}
	}

	notification := &models.Notification{
		UserID:  user.ID,
		Active:  false,
		Suggest: true,
	}
	if err := tx.WithContext(ctx).Create(notification).Error; err != nil {
		tx.Rollback()
		return err
	}

	code, err := hash.GenerateEmailVerification(client, user.Email, rctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := mail.SendRegistrationEmail(user.Email, code, dialer); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
