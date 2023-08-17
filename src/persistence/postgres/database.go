package postgres

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Error In Database Connection")
	}

	Database = db

}
