package model

import "time"

type VaccineCertificate struct {
	VaccineCertificateId      uint          `gorm:"primaryKey" json:"vaccine_certificate_id"`
	VaccineName               Vaccine       `json:"vaccine_name"`
	PatientId                 Patient       `json:"patient_id"`
	NurseId                   Nurse         `json:"nurse_id"` // aply the associations concepts
	EstablishmentName         Establishment `json:"establishment_name"`
	VaccineAplicationDate     time.Time     `json:"vaccine_aplication_date"`
	NextVaccineAplicationDate time.Time     `json:"next_vaccine_aplication_date"`
}
