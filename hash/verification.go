package hash

import (
	"context"
	"errors"
	"fmt"
	"time"

	"git.sindadsec.ir/asm/backend/utils"
	"github.com/redis/go-redis/v9"
)

func GenerateEmailVerification(client *redis.Client, email string, rctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(rctx, time.Second*5)
	defer cancel()

	randCode := utils.GenerateRandomCode()
	_, err := client.Set(ctx, fmt.Sprintf("verification-%s", email), randCode, time.Second*120).Result()
	if err != nil {
		return "", err
	}

	return randCode, nil
}

func GetEmailVerificationByEmail(client *redis.Client, email string, rctx context.Context) error {
	ctx, cancel := context.WithTimeout(rctx, time.Second*5)
	defer cancel()

	_, err := client.Get(ctx, fmt.Sprintf("verification-%s", email)).Result()
	if err != nil {
		switch {
		case errors.Is(err, redis.Nil):
			return ErrNilKey
		default:
			return err
		}
	}

	return nil
}
