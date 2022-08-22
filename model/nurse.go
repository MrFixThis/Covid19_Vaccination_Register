package model

type Nurse struct {
	Person
	NurseRol            string               `json:"nurse_rol"`
	VaccineCertificates []VaccineCertificate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vaccine_certificates"`
}
