package model

type Nurse struct {
	Person
	NurseRol           string               `json:"nurse_rol"`
	VaccineCertificate []VaccineCertificate `json:"vaccine_certificates"`
}
