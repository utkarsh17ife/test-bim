package models

import "time"

type Tasks struct {
	MongoID         string
	TaskID          string
	TaskName        string
	TaskDescription string
	TaskOwner       string
	AssignedTo      string
	AssignedBy      string
	StartDate       time.Time
	EndDate         time.Time
	Status          string
	UpdatedAt       string
	CreatedAt       string
}
