package models

import "time"

type Todo struct {
	TodoID      int       `json:"todo_id"`
	Content     string    `json:"content"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedFlag bool      `json:"deleted_flag"`
}
