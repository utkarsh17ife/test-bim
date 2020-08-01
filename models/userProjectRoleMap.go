package models

import "time"

type UserProjectRoleMap struct {
	MongoID           string
	UserProjectRoleID string
	ProjectID         string
	UserID            string
	RoleID            string
	UpdatedAt         time.Time
	CreatedAt         time.Time
}
