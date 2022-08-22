package database

import (
	"github.com/Covid19_Vaccination_Register/model"
	"gorm.io/gorm"
)

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(
		&model.Address{},
		&model.Address{},
		&model.Vaccine{},
		&model.Patient{},
		&model.Nurse{},
		&model.Establishment{},
		&model.Administrator{},
		&model.VaccineCertificate{},
	)
}
