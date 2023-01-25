package db

import (
	"fmt"
	stdlog "log"
	"time"

	"github.com/joint-online-judge/go-horse/app/querys"
	"github.com/joint-online-judge/go-horse/pkg/configs"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB gorm connector
var DB *gorm.DB

func ConnectPostgres() {
	var err error // define error here to prevent overshadowing the global DB
	conf := configs.Conf
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
	querys.DB = DB
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %+v", err)
	}
	// TODO: run auto migrate
	// err = DB.AutoMigrate(&model.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
