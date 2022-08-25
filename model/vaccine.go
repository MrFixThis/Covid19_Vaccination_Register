package model

import (
	"time"

	"gorm.io/gorm"
)

type Vaccine struct {
	gorm.Model
	Name                string               `json:"name"`
	ManufactureDate     time.Time            `json:"manufacture_date"`
	ExpirationDate      time.Time            `json:"expiration_date"`
	VaccineCertificates []VaccineCertificate `json:"vaccine_certificates"`
}
