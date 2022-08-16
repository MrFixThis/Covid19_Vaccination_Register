package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	City             string          `json:"city"`
	NeighborhoodName string          `json:"neighborhood_name"`
	StreetName       string          `json:"street_name"`
	StreetNumber     uint32          `json:"street_number"`
	CardinalLocation string          `json:"cardinal_location"`
	Patients         []Patient       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"patients"`
	Establishments   []Establishment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"establishments"`
}
