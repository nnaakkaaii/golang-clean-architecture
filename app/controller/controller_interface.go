package controller

import (
	"github.com/labstack/echo"
)

type UserControllerInterface interface {
	ListUsers(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	RetrieveUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}
