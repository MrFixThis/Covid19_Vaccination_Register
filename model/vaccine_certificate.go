package model

import "time"

type VaccineCertificate struct {
	VaccineCertificateId      uint          `gorm:"primaryKey" json:"vaccine_certificate_id"`
	VaccineName               string        `json:"vaccine_name"`
	PatientId                 Patient       `gorm:"embedded" json:"patient_id"`
	NurseId                   Nurse         `gorm:"embedded" json:"nurse_id"`
	EstablishmentName         Establishment `gorm:"embedded" json:"establishment_name"`
	VaccineAplicationDate     time.Time     `json:"vaccine_aplication_date"`
	NextVaccineAplicationDate time.Time     `json:"next_vaccine_aplication_date"`
} // TODO: implement hooks for manage the transactions
