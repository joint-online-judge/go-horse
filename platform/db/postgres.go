package db

import (
	"fmt"
	"time"

	"github.com/joint-online-judge/go-horse/app/querys"
	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
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
	logLevel := gorm_logger.Silent
	if conf.DBEcho {
		logLevel = gorm_logger.Info
	}
	newLogger := gorm_logger.New(
		logrus.StandardLogger(),
		gorm_logger.Config{
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
		logrus.Fatalf("failed to connect to Postgres: %+v", err)
	}
	// TODO: run auto migrate
	// err = DB.AutoMigrate(&model.User{})
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
}
