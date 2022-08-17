package database

import (
	"github.com/Covid19_Vaccination_Register/model"
	"gorm.io/gorm"
)

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(
		&model.Address{},
		&model.Establishment{},
		&model.Patient{},
		&model.Nurse{},
		&model.Vaccine{},
		&model.Administrator{},
		&model.VaccineCertificate{},
	)
}
