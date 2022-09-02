package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	City             string        `json:"city"`
	NeighborhoodName string        `json:"neighborhood_name"`
	StreetName       string        `json:"street_name"`
	StreetNumber     string        `json:"street_number"`
	CardinalLocation string        `json:"cardinal_location"`
	PatientID        uint          `json:"patient_id"`
	Establishment    Establishment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"establishment"`
}
