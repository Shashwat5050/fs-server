package models

import (
	"time"
)

type User struct {
	ID             string     `db:"id" json:"id"`
	Email          string     `db:"email" json:"email"`
	Password       string     `db:"password" json:"password"`
	IsVerified     bool       `db:"is_verified" json:"is_verified"`
	IsTwoFaEnabled bool       `db:"is_two_fa_enabled" json:"is_two_fa_enabled"`
	IsActive       bool       `db:"is_active" json:"is_active"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdateAt       *time.Time `db:"updated_at" json:"updated_at"`
	DeleteAt       *time.Time `db:"deleted_at" json:"deleted_at"`
}
