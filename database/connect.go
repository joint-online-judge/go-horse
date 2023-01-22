package database

import (
	"log"
	"os"

	"github.com/joint-online-judge/go-horse/dal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	env := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dal.SetDefault(DB)
	// TODO: run auto migrate
	// err = DB.AutoMigrate(&model.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
