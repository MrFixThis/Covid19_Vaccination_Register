package model

import (
	"time"

	"gorm.io/gorm"
)

type VaccineCertificate struct {
	gorm.Model
	VaccineName                string    `json:"vaccine_name"`
	PatientID                  uint      `json:"patient_id"`
	NurseID                    uint      `json:"nurse_id"`
	EstablishmentName          string    `json:"establishment_name"`
	VaccineApplicationDate     time.Time `json:"vaccine_application_date"`
	NextVaccineApplicationDate time.Time `json:"next_vaccine_application_date"`
}
