package database

import (
	"github.com/Covid19_Vaccination_Register/model"
)

func migrateSchema() {
	DB.AutoMigrate(
		&model.Address{},
		&model.Vaccine{},
		&model.Patient{},
		&model.Nurse{},
		&model.Establishment{},
		&model.VaccineCertificate{},
		&model.Administrator{},
	)
}
