package notify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"iceline-hosting.com/core/models"
)

var (
	ErrInvalidUserID = errors.New("invalid user id")
)

type queuer interface {
	Push(ctx context.Context, key string, value ...interface{}) error
}

type client struct {
	kvClient queuer
}

func NewClient(kvClient queuer) client {
	return client{
		kvClient: kvClient,
	}
}

func (c client) PushNotification(ctx context.Context, notification models.Notification) error {
	userID, ok := ctx.Value(models.KeyUserID).(string)
	if !ok {
		return fmt.Errorf("%w: the type must be string", ErrInvalidUserID)
	}
	if userID == "" {
		return fmt.Errorf("%w: userID must be set in context", ErrInvalidUserID)
	}

	notification.ID = uuid.New().String()
	notification.ToUserID = userID
	notification.CreatedAt = time.Now()

	value, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	return c.kvClient.Push(ctx, models.NotificationQueue, string(value))
}
