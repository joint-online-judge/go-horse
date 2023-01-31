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

func createDatabase(gormConfig *gorm.Config) error {
	conf := configs.Conf
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUsername,
		conf.DBPassword,
		"postgres",
	)
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return err
	}
	// check if db exists
	stmt := fmt.Sprintf(
		"SELECT * FROM pg_database WHERE datname = '%s';",
		conf.DBName,
	)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		return rs.Error
	}
	// if not create it
	rec := make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", conf.DBName)
		if rs := db.Exec(stmt); rs.Error != nil {
			return rs.Error
		}
		// close db connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func connectDatabase(gormConfig *gorm.Config) (err error) {
	conf := configs.Conf
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUsername,
		conf.DBPassword,
		conf.DBName,
	)
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	querys.DB = DB
	return
}

func migrateDatabase() error {
	return DB.AutoMigrate(
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
}

func getGormConfig() gorm.Config {
	conf := configs.Conf
	logLevel := gorm_logger.Silent
	if conf.DBEcho {
		logLevel = gorm_logger.Info
	}
	gormConfig := gorm.Config{Logger: gorm_logger.New(
		logrus.StandardLogger(),
		gorm_logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)}
	return gormConfig
}

func ConnectPostgres() {
	gormConfig := getGormConfig()
	if err := createDatabase(&gormConfig); err != nil {
		logrus.Fatalf("failed to create db in Postgres: %+v", err)
	}
	if err := connectDatabase(&gormConfig); err != nil {
		logrus.Fatalf("failed to connect to Postgres: %+v", err)
	}
	if err := migrateDatabase(); err != nil {
		logrus.Fatalf("failed to migrate in Postgres: %+v", err)
	}
}
