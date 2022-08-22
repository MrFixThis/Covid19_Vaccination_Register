package model

import "gorm.io/gorm"

type Establishment struct {
	gorm.Model
	Name                string               `json:"name"`
	AddressID           uint                 `json:"address_id"`
	ContactNumber       string               `json:"contact_number"`
	ContactEmailAddress string               `json:"contact_email_address"`
	VaccineCertificates []VaccineCertificate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vaccine_certificates"`
}
