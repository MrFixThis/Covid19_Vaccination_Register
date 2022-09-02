package model

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName                    string    `json:"first_name"`
	LastName                     string    `json:"last_name"`
	IdentificationNumber         string    `json:"identification_number"`
	IdentificationExpeditionDate time.Time `json:"identification_expedition_date"`
	PhoneNumber                  string    `json:"phone_number"`
	EmailAddress                 string    `json:"email_address"`
}
