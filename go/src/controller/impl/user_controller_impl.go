package controller

import (
	"main/dto"
	store "main/store/interface"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserStore store.UserStore
}

func (f *UserControllerImpl) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, f.UserStore.GetAllUser())
}

func (f *UserControllerImpl) CreateUser(c echo.Context) error {
	var userDto dto.CreateUserDto
	if e := c.Bind(userDto); e != nil {
		return e
	}

	u := f.UserStore.CreateUser(userDto)
	return c.JSON(http.StatusOK, u)
}