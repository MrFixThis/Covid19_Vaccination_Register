package model

// NurseRole type refers to the unique role that a nurse must have
type NurseRole int

// Nurse roles enum
const (
	Auxiliar NurseRole = iota
	Vaccinator
	Boss
)

type Nurse struct {
	Person
	NurseRole           NurseRole            `json:"nurse_role"`
	VaccineCertificates []VaccineCertificate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vaccine_certificates"`
}
