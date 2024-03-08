package models

import "time"

type BackupInfo struct {
	ID         string     `db:"id"`
	ServerName string     `db:"server_name"`
	Filename   string     `db:"filename"`
	Size       uint64     `db:"size"`
	Locked     bool       `db:"locked"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}
