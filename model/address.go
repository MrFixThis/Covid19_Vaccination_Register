package model

type Address struct {
	AddressID        uint   `gorm:"primarykey" json:"address_id"`
	Operation
	City             string `json:"city"`
	NeighborhoodName string `json:"neighborhood_name"`
	StreetName       string `json:"street_name"`
	StreetNumber     uint32 `json:"street_number"`
	CardinalLocation string `json:"cardinal_location"`
}
