package models

import "time"

type Company struct {
	MongoID              string
	CompanyID            string
	CompanyName          string
	OwnerUserID          string
	RegistrationDate     time.Time
	Location             string
	Address              string
	CompanyEmail         string
	CompanyContactNumber string
	WebsiteLink          string
	UpdatedAt            time.Time
	CreatedAt            time.Time
}
