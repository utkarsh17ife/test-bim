package models

import "time"

type Comment struct {
	MongoID    string
	CommentID  string
	ProjectID  string
	TaskID     string
	PostedByID string
	Text       string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}
