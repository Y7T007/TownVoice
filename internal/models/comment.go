package models

import (
	"time"
)

type Comment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	EntityID  string    `json:"entity_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
