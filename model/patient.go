package model

type Patient struct {
	Person
	Addresses          []Address          `json:"addresses"`
	VaccineCertificate VaccineCertificate `json:"vaccine_certificate"`
}
