package controller

import (
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
