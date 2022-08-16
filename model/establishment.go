package model

type Establishment struct {
	Name                string               `gorm:"primarykey" json:"name"`
	Operation
	AddressID           uint                 `json:"address_id"`
	ContactNumber       string               `json:"contact_number"`
	ContactEmailAddress string               `json:"contact_email_address"`
	VaccineCertificates []VaccineCertificate `json:"vaccine_certificates"`
}
