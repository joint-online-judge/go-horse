package db

import (
	"fmt"
	"time"

	"github.com/joint-online-judge/go-horse/app/models"
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
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUsername,
		conf.DBPassword,
		conf.DBName,
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
	gormConfig := gorm.Config{Logger: newLogger}
	DB, err = gorm.Open(postgres.Open(dsn), &gormConfig)
	querys.DB = DB
	// TODO: create the database if non-exist
	if err != nil {
		logrus.Fatalf("failed to connect to Postgres: %+v", err)
	}
	err = DB.AutoMigrate(
		&models.DomainInvitation{},
		&models.DomainRole{},
		&models.DomainUser{},
		&models.Domain{},
		&models.ProblemConfig{},
		&models.ProblemGroup{},
		&models.ProblemProblemSetLink{},
		&models.ProblemSet{},
		&models.Problem{},
		&models.Record{},
		&models.UserLatestRecord{},
		&models.UserOauthAccount{},
		&models.User{},
	)
	if err != nil {
		logrus.Fatal(err)
	}
}
