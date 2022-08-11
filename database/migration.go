package database

import (
	"github.com/Covid19_Vaccination_Register/model"
	"gorm.io/gorm"
)

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(
		&model.Patient{},
		&model.Nurse{},
		&model.Establishment{},
		&model.Address{},
		&model.Vaccine{},
		&model.VaccineCertificate{},
	)
} // read the doc to see if the are other ways to migrate the schema
