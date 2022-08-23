package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:@BrysDB130502.@tcp(127.0.0.1:3306)/covid19_vaccine_register?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB
var err error

func InitializeDatabase() {
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // <- read more about this
	})
	if err != nil {
		panic(err.Error())
	}

	// To migrate the schema
	migrateSchema()
}
