package model

type Nurse struct {
	Person
	NurseRol            string               `json:"nurse_rol"`
	VaccineCertificates []VaccineCertificate `json:"vaccine_certificates"`
}
