package model

type Patient struct {
	Person
	AddressID           uint                 `json:"address_id"`
	VaccineCertificates []VaccineCertificate `json:"vaccine_certificates"`
}
