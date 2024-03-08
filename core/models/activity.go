package models

import "time"

type GsActivity struct {
	ServerName string `json:"server_name" db:"server_name" binding:"required"`
	UserActivity
}

type UserActivity struct {
	UserID       string    `json:"user_id" db:"user_id" binding:"required"`
	UserEmail    string    `json:"user_email" db:"user_email" binding:"required"`
	ActivityType string    `json:"activity_type" db:"activity_type" binding:"required"`
	Time         time.Time `json:"time" db:"time"`
}
