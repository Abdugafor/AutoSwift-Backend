package db

import (
	"AutoSwift/Models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

const dbCredentials = `host=localhost user=postgres password=1234 dbname=AutoSwift port=5432 sslmode=disable `

func StartDatabase() {
	var err error
	db, err = gorm.Open(postgres.Open(dbCredentials), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Println("db.setupPostgresProduction err:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(100)
	autoMigrate()
	log.Println("DB successfully connected! ")
}

func autoMigrate() {
	for _, model := range []interface{}{
		(*Models.Car)(nil),
	} {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Println("create model %s : %s", model, err)
		}
	}
}

func GetDb() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	sqlDB.Close()

	if err != nil {
		fmt.Println("Error on closing the DB: ", err)
	}
	fmt.Println("DB closed!")
}
