package models

import "time"

type Type int

const (
	_ Type = iota
	ServerNotify
	WelcomeNotify
	ResetPasswordNotify

	NotificationQueue = "notification_queue"
	KeyUserID         = "userID"
)

func (t Type) String() string {
	switch t {
	case ServerNotify:
		return "server"
	case WelcomeNotify:
		return "welcome"
	case ResetPasswordNotify:
		return "reset_password"
	default:
		return "unknown"
	}
}

type Notification struct {
	ID        string
	Type      Type
	ToUserID  string
	Values    map[string]interface{}
	CreatedAt time.Time
	SentAt    time.Time
}
