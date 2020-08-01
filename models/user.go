package models

import "time"

// user model
type User struct {
	MongoID       string // mongo auto id
	UserID        string // we will generate this eg. U10001, U10002 [PK]
	Name          string
	Password      string
	Email         string // unique
	ContactNumber string // unique
	CompanyUserID string // user id in his company, can be null
	CompanyID     string // company id
	UpdatedAt     time.Time
	CreatedAt     time.Time
}
