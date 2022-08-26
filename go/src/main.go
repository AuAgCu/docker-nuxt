package main

import (
	"main/config"
	controller "main/controller/impl"
	"main/entity"
	db "main/repository"
	repository "main/repository/impl"
	service "main/service/impl"
	store "main/store/impl"
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
	// TODO: ここの初期化別の場所でやりたい、di containerとか使うのがいいかな？
	userController := controller.UserControllerImpl{
		UserStore: &store.UserStoreImpl{
			UserService: &service.UserServiceImpl{
				UserRepository: &repository.UserRepositoryImpl{Db: db.DB},
			},
		},
	}

	e.GET("/", articleIndex)
	e.GET("/api/hello", hello)
	e.GET("/api/users", userController.GetUsers)

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
