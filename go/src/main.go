package main

import (
	"main/config"
	container "main/container"
	controller "main/controller"
	db "main/repository"
	repository "main/repository/impl"
	service "main/service/impl"
	store "main/store/impl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "gorm.io/gorm"
)

var e = createMux()

func main() {
	initDiContainer()
	// TODO: ここの初期化別の場所でやりたい、di containerとか使うのがいいかな？
	userController := controller.UserControllerImpl{
		UserStore: &store.UserStoreImpl{
			UserService: &service.UserServiceImpl{
				UserRepository: &repository.UserRepositoryImpl{Db: db.DB},
			},
		},
	}

	e.GET("/api/hello", userController.GetUsers)
	e.POST("/api/user", userController.CreateUser)

	e.Logger.Fatal(e.Start(":3001"))
}

func initDiContainer() *container.DiContainer {
	container := container.NewDiContainer()
	container.Register(controller.NewUserController)

	return container
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
