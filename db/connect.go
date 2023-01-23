package db

import (
	"fmt"
	stdlog "log"
	"time"

	"github.com/joint-online-judge/go-horse/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB
	conf := config.Config
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.DBHost,
		conf.DBUsername,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
	)
	logLevel := logger.Silent
	if conf.DBEcho {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		stdlog.New(log.StandardLogger().Out, "\r\n", stdlog.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	// TODO: run auto migrate
	// err = DB.AutoMigrate(&model.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
