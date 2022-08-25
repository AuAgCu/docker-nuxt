package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var e = createMux()

func main() {
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	m, err := migrate.New(
		"file://migration",
		"postgres://postgres:password@db:5432/docker-nuxt?sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
