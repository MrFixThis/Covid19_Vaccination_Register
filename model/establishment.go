package model

type Establishment struct {
	Name                string `gorm:"primarykey" json:"name"`
	Operation
	Adress              Address `gorm:"embedded;column:establishment_adress" json:"establishment_adress"`
	ContactNumber       uint16  `json:"contact_number"`
	ContactEmailAddress string  `json:"contact_email_address"`
}
