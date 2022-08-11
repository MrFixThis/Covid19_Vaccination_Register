package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func ConnectDatabase() *gorm.DB {
	const dsn = `root:@BrysDB130502.@tcp(127.0.0.1:3306)/covid19_vaccine_
register_db?charset=utf8mb4&parseTime=True&loc=Local`

	// read the documentation to see the advanced db's connection configuration
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	return db
}
