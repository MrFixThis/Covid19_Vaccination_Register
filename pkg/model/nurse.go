package model

// NurseRole type refers to the unique role that a nurse must have
// 1: Auxiliar
// 2: Vaccinator
// 3: Boss
type NurseRole int

type Nurse struct {
	Person
	NurseRole           NurseRole            `json:"nurse_role"`
	VaccineCertificates []VaccineCertificate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vaccine_certificates"`
}
