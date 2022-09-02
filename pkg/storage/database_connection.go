package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error
const dsn = "root:@BrysDB130502.@tcp(127.0.0.1:3306)" +
	"/covid19_vaccine_register?charset=utf8mb4&parseTime=True&loc=Local"

func InitializeDatabase() {
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true, // <- read more about this
	})
	if err != nil {
		panic(err.Error())
	}

	// To migrate the schema
	migrateSchema()
}
