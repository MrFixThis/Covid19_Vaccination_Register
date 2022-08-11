package model

import "time"

type Vaccine struct {
	VaccineName     string         `gorm:"primaryKey" json:"vaccine_name"`
	Operation
	ManufactureDate time.Time      `json:"manufacture_date"`
	ExpirationDate  time.Time      `json:"expiration_date"`
}
