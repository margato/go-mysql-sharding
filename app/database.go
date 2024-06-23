package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func createGormDatabaseConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panicf("Could not establish connection to database: %s\n", err)
	}

	return db
}

func CreateDatabaseShardsConnections() map[int]*gorm.DB {
	return map[int]*gorm.DB{
		0: createGormDatabaseConnection(
			os.Getenv("SHARD_0_DSN"),
		),
		1: createGormDatabaseConnection(
			os.Getenv("SHARD_1_DSN"),
		),
		2: createGormDatabaseConnection(
			os.Getenv("SHARD_2_DSN"),
		),
	}
}
