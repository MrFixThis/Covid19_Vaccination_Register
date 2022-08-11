package model

type Establishment struct {
	EstablishmentName string `gorm:"primarykey" json:"establishment_name"`
	Operation
	Adress              Address `gorm:"embedded;column:establishment_adress" json:"establishment_adress"`
	ContactNumber       uint16  `json:"contact_number"`
	ContactEmailAddress string  `json:"contact_email_address"`
}
