package models

import "time"

type Document struct {
	MongoID         string
	DocumentID      string
	ProjectID       string
	TaskID          string
	UplodedByUserID string
	DocumentURL     string
	DocumentTypeID  string
	VisibilityType  string // task_based, project_based
	UpdateAt        time.Time
	CreatedAt       time.Time
}
