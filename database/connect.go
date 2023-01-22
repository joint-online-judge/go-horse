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
	db, err = gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dal.SetDefault(db)
	// TODO: run auto migrate
	// err = db.AutoMigrate(&model.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
