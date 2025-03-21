package clients

import (
	"context"
	"net/http"
	"time"

	"git.sindadsec.ir/asm/backend/utils"
)

type disposablePaylaod struct {
	Disposable *string `json:"disposable" validate:"oneof=true false"`
}

func IsDisposable(api, email string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var payload disposablePaylaod
	if err := utils.ReadJson(res.Body, &payload); err != nil {
		return nil, err
	}

	return payload.Disposable, nil
}

func CheckDisposableApiHealth(api string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	return err
}
