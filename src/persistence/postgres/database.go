package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {

	url := "host=localhost user=postgres password=12345678 dbname=go-bank port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("Error In Database Connection")
	}

	Database = db

}
