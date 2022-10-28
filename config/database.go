package config

import (
	"fmt"
	"github.com/vodeacloud/hr-api/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {
	if db != nil {
		return db
	}

	gormConfig := &gorm.Config{}
	if !IsProd() {
		gormConfig.Logger = gormlogger.Default.LogMode(gormlogger.Info)
	}

	dbCon, err := gorm.Open(mysql.Open(GetDatabaseDSN()), gormConfig)
	if err != nil {
		logger.Fatalf("failed create connection to database: %v", err)
		return nil
	}

	sqlDB, _ := dbCon.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = dbCon
	return db
}

func GetDatabaseDSN() string {
	config := GetConfig()
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBDatabase,
	)
}

func CloseDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		logger.Fatalf("failed close connection database: %v", err)
		return
	}
	if err = dbSQL.Close(); err != nil {
		logger.Fatalf("failed close connection database: %v", err)
	}
}
