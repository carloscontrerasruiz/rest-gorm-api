package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user= user password=pass dbname=mydb port=5432"
var DB *gorm.DB

func DbConnection() {
	var err error
	DB, err = gorm.Open(
		postgres.Open(DSN),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connected")
}
