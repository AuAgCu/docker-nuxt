package main

import (
	"main/config"
	"main/entity"
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
	e.GET("/api/hello", hello)

	e.Logger.Fatal(e.Start(":3001"))
}

func createMux() *echo.Echo {
	m, err := migrate.New(
		"file://migration",
		config.DB_URL,
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

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, entity.User{FirstName: "first", LastName: "last"})
}
