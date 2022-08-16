package model

import "time"

type Vaccine struct {
	Name               string                `gorm:"primaryKey" json:"name"`
	Operation
	ManufactureDate    time.Time             `json:"manufacture_date"`
	ExpirationDate     time.Time             `json:"expiration_date"`
	VaccineCertificates []VaccineCertificate `json:"vaccine_certificates"`
}
