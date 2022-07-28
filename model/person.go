package model

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Identification uint32 `json:"identification"`
	PhoneNumber    uint32 `json:"phone_number"`
	EmailAddress   string `json:"email_address"`
}
