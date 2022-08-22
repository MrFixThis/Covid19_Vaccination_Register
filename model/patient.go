package model

type Patient struct {
	Person
	Addresses          []Address          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"address_id"`
	VaccineCertificate VaccineCertificate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vaccine_certificate"`
}
