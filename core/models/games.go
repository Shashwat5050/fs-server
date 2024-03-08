package models

import (
	"time"

	"github.com/lib/pq"
)

type Game struct {
	ID          string         `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description string         `db:"description" json:"description"`
	Image       string         `db:"image" json:"image"`
	Envs        pq.StringArray `db:"envs" json:"envs"`
	Ports       pq.Int32Array  `db:"ports" json:"ports"`
	Volumes     pq.StringArray `db:"volumes" json:"volumes"`
	CPU         int            `db:"cpu" json:"cpu"`
	Memory      int            `db:"memory" json:"memory"`
	Command     string         `db:"command" json:"command"`
	Args        pq.StringArray `db:"args" json:"args"`
	WithDB      bool           `db:"with_db" json:"with_db"`
	CoverImage  string         `db:"cover_image" json:"cover_image"`
	GameModID   int            `db:"game_mod_id" json:"game_mod_id"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time     `db:"updated_at" json:"updated_at"`
}
