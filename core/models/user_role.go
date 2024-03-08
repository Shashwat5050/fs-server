package models

type UserRole struct {
	ID       int64  `db:"id" json:"id"`
	UserID   string `db:"user_id" json:"user_id"`
	Role     string `db:"role" json:"role"`
	Resource string `db:"resource" json:"resource"`
}
