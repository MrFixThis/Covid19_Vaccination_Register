package model

type Patient struct {
	Person
	VaccineCertificateID uint
	VaccineCertificate   VaccineCertificate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ResidenceAddress     Address            `json:"residence_address"`
}

// one-to-one -> each instance of this model 'belongs to' each instance of
// vaccine_certificate model

// one-to-many -> each
