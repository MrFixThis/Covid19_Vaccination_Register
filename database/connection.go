package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const dsn = `root:@BrysDB130502.@tcp(127.0.0.1:3306)/covid19_vaccine_register
_db?charset=utf8mb4&parseTime=True&loc=Local`

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // for default conf
	// until we read the documentation
	if err != nil {
		panic("could not connect to the database")
	}

	return db
} // TODO: Create the database here in ubuntu
