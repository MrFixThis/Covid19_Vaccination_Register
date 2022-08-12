package model

import (
	"time"

	"gorm.io/gorm"
)

type VaccineCertificate struct {
	gorm.Model
	VaccineName               Vaccine       `json:"vaccine_name"`
	NurseId                   Nurse         `json:"nurse_id"`
	EstablishmentName         Establishment `json:"establishment_name"`
	VaccineAplicationDate     time.Time     `json:"vaccine_aplication_date"`
	NextVaccineAplicationDate time.Time     `json:"next_vaccine_aplication_date"`
}
