package model

type Patient struct {
	Person
	AddressID           uint               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"address_id"`
	VaccineCertificate  VaccineCertificate `json:"vaccine_certificate"`
}
