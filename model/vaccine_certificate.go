package model

import (
	"time"

	"gorm.io/gorm"
)

type VaccineCertificate struct {
	gorm.Model
	VaccineID                 uint      `json:"vaccine_id"`
	PatientID                 uint      `json:"patient_id"`
	NurseID                   uint      `json:"nurse_id"`
	EstablishmentName         string    `json:"establishment_name"`
	VaccineAplicationDate     time.Time `json:"vaccine_aplication_date"`
	NextVaccineAplicationDate time.Time `json:"next_vaccine_aplication_date"`
}
