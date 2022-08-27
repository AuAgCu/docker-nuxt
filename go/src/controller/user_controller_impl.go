package controller

import (
	"log"
	"main/dto"
	store "main/store/interface"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserStore store.UserStore
}

func NewUserController(userStore store.UserStore) UserController {
	return &UserControllerImpl{UserStore: userStore}
}

func (f *UserControllerImpl) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, f.UserStore.GetAllUser())
}

func (f *UserControllerImpl) CreateUser(c echo.Context) error {
	var userDto dto.CreateUserDto
	if e := c.Bind(&userDto); e != nil {
		log.Println(e)
		return e
	}

	u := f.UserStore.CreateUser(userDto)
	return c.JSON(http.StatusOK, u)
}
