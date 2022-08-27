package main

import (
	"main/config"
	container "main/container"
	co "main/controller"
	repository "main/repository"
	service "main/service"
	store "main/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var e = createMux()

func main() {
	repository.InitDb(config.DB_URL, &gorm.Config{})
	c := initDiContainer()

	uc := c.GetInstance(func(arg co.UserController) {}).Interface().(co.UserController)

	e.GET("/api/user", uc.GetUsers)
	e.POST("/api/user", uc.CreateUser)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}

func initDiContainer() *container.DiContainer {
	c := container.NewDiContainer()
	c.Register(co.NewUserController)
	c.Register(store.NewUserStore)
	c.Register(service.NewUserService)
	c.Register(repository.NewUserRepository)

	return c
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
