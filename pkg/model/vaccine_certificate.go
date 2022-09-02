package model

import (
	"time"

	"gorm.io/gorm"
)

type VaccineCertificate struct {
	gorm.Model
	VaccineID                  uint      `json:"vaccine_id"`
	PatientID                  uint      `json:"patient_id"`
	NurseID                    uint      `json:"nurse_id"`
	EstablishmentID            uint      `json:"establishment_id"`
	VaccineApplicationDate     time.Time `json:"vaccine_application_date"`
	NextVaccineApplicationDate time.Time `json:"next_vaccine_application_date"`
}
