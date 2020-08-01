package models

import "time"

type Project struct {
	MongoID      string // mongo auto id
	ProjectID    string
	ProjectName  string
	CompanyID    string
	Status       string // enum
	Location     string
	StartDate    time.Time
	EndDate      time.Time
	EstimateCost int64
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
