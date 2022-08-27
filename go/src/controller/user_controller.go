package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
}
