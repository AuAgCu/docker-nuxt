package repository

import (
	"main/config"

	"github.com/jinzhu/gorm"
)

var DB = gormConnect()

func gormConnect() *gorm.DB {
	DBMS := "postgres"
	CONNECT := config.DB_URL
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}
