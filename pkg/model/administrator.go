package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Administrator struct {
	gorm.Model
	AdminName string `gorm:"unique" json:"admin_name"`
	Password  string `gorm:"size:255" json:"password,omitempty"`
	IsMaster  bool   `json:"-"`
}

// BeforeSave hook encrypts the administrator's password before saving it in
// the database. If the administrator alredy exists in the database but its
// fields has a new value, then an update will be made, encripting the new
// password if it was changed
func (a *Administrator) BeforeSave(tx *gorm.DB) (err error) {
	var ar Administrator
	r := tx.First(&ar, a.ID).Error

	if r != nil || (ar.Password != "" && ar.Password != a.Password) {
		h, err := encryptPassword(a.Password)
		if err != nil {
			return err
		}
		a.Password = h
	}
	return
}

// encryptPassword encrypts a given string password and returns its hash and
// a possible error result
func encryptPassword(rp string) (h string, err error) {
	hb, err := bcrypt.GenerateFromPassword([]byte(rp), bcrypt.DefaultCost)
	return string(hb), err
}
