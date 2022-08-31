package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Administrator struct {
	gorm.Model
	AdminName string `gorm:"unique" json:"admin_name"`
	Password  string `gorm:"size:255" json:"password,omitempty"`
	IsMaster  bool   `json:"-"`
}

func (a *Administrator) BeforeSave(tx *gorm.DB) (err error) {
	tx = tx.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
	var ar Administrator
	r := tx.First(&ar, a.ID).Error

	switch {
	case r != nil, (ar.Password != "" && a.Password != ar.Password):
		h, err := encryptPassword(a.Password)
		if err != nil {
			return err
		}
		a.Password = h
	}
	return
}

// encryptPassword encrypts a given string password and returns its hash and
// a possible error value
func encryptPassword(rp string) (h string, err error) {
	hb, err := bcrypt.GenerateFromPassword([]byte(rp), bcrypt.DefaultCost)
	return string(hb), err
}
