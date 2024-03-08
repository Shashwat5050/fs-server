package models

import (
	"time"

	"github.com/lib/pq"
)

type Sftp struct {
	ID         string      `db:"id"`
	ServerName string      `db:"server_name"`
	Username   string      `db:"username"`
	Port       int         `db:"port"`
	CreatedAt  time.Time   `db:"created_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
	DeletedAt  pq.NullTime `db:"deleted_at"`
}
