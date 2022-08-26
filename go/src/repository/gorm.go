package repository

import (
	"main/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = gormConnect()

func gormConnect() *gorm.DB {
	CONNECT := config.DB_URL
	db, err := gorm.Open(postgres.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
