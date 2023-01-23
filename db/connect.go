package db

import (
	"log"

	"github.com/joint-online-judge/go-horse/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB
	dsn := config.Config.DatabaseUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// TODO: run auto migrate
	// err = DB.AutoMigrate(&model.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
