package models

import (
	"time"
)

type Task struct {
	ID                string     `json:"id" db:"id"`
	ScheduleID        string     `json:"schedule_id" db:"schedule_id"`
	Action            string     `json:"action" db:"action"`
	TimeOffset        int        `json:"time_offset" db:"time_offset"`
	Payload           string     `json:"payload" db:"payload"`
	ContinueOnFailure bool       `json:"continue_on_failure" db:"continue_on_failure"`
	Status			  string	 `db:"status"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         *time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func(task *Task)TableName()string{
	return "task"
}
