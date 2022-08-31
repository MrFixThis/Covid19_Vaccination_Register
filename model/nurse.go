package model

type Nurse struct {
	Person
	NurseRole           string               `json:"nurse_role"`
	VaccineCertificates []VaccineCertificate `json:"vaccine_certificates"`
}
