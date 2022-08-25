package config

import (
	"os"
)

var (
	POSTGRES_USER     = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB       = os.Getenv("POSTGRES_DB")
	POSTGRES_PORT     = os.Getenv("POSTGRES_PORT")
)
