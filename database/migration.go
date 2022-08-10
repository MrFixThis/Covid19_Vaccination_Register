package database

import (
	"github.com/Covid19_Vaccination_Register/model"
	"gorm.io/gorm"
)

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(
		&model.Patient{},
		&model.Nurse{},
		&model.Establishment{}, // so, we need to read the other way to
		&model.Address{},      // migrate the schema. See the gorm doc
		&model.Vaccine{},
		&model.VaccineCertificate{},
	)
}
