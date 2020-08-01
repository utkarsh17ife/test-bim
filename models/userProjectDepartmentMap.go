package models

import "time"

type UserProjectDepartmentMap struct {
	MongoID                 string
	UserProjectDepartmentID string
	ProjectID               string
	UserID                  string
	DepartmentID            string
	UpdatedAt               time.Time
	CreatedAt               time.Time
}
