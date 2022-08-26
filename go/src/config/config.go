package config

import (
	"fmt"
	"os"
)

var (
	POSTGRES_USER     = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB       = os.Getenv("POSTGRES_DB")
	POSTGRES_PORT     = os.Getenv("POSTGRES_PORT")
	POSTGRES_HOST     = os.Getenv("POSGTRES_HOST")

	DB_URL = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB)
)
