package model

type Address struct {
	AddressID        uint   `gorm:"primarykey" json:"address_id"`
	City             string `json:"city"`
	Neighborhood     string `json:"neighborhood"`
	Street           string `json:"street"`
	StreetNumber     uint32 `json:"street_number"`
	CardinalLocation string `json:"cardinal_location"`
}
