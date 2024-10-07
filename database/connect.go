package database

import (
	"gobackend/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("./arabic.db"), &gorm.Config{})

	if err != nil {
		log.Panicln("Could not connect to the database. " + err.Error())
	}

	DB = db
	log.Println("Connect to postgres database successfully!")
	db.AutoMigrate(
		&models.User{},
		&models.Topic{},
		&models.Word{},
	)

}
