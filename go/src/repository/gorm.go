package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(dbUrl string, config *gorm.Config) {
	DB = gormConnect(dbUrl, config)
}

func gormConnect(dbUrl string, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), config)
	if err != nil {
		panic(err.Error())
	}

	return db
}
