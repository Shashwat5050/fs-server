package models

import "time"

type Schedule struct {
	ID               string     `json:"id" db:"id"`
	ServerID         string     `json:"server_id" db:"server_id"`
	ScheduleName     string     `json:"schedule_name" db:"schedule_name"`
	Frequency        string     `json:"frequency" db:"frequency"`
	CustomExpression string     `json:"custom_expression" db:"custom_expression"`
	ServerOnline     bool       `json:"server_online" db:"server_online"`
	ScheduleEnabled  bool       `json:"schedule_enabled" db:"schedule_enabled"`
	TimeZone         string     `json:"time_zone" db:"time_zone"`
	LastStatus       string		`db:"last_status"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        *time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
