package model

type Patient struct {
	Person
	ResidenceAddress Address `json:"residence_address"`
}
