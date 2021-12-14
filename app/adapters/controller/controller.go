package controller

import (
	"clean_architecture/app/usecases/interactor"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type UserController struct {
	interactor *interactor.UserInteractor
}

func NewUserController(interactor *interactor.UserInteractor) *UserController {
	return &UserController{interactor: interactor}
}

func (controller *UserController) ListUsers(ctx echo.Context) error {
	res, err := controller.interactor.ListUsers()
	if err != nil {
		return handle(ctx, err)
	}
	return ctx.JSONBlob(http.StatusOK, res)
}

func (controller *UserController) CreateUser(ctx echo.Context) error {
	var r requestBody
	if err := ctx.Bind(&r); err != nil {
		return handle(ctx, fmt.Errorf("入力をパースできません"))
	}
	obj := r.convert()
	res, err := controller.interactor.CreateUser(obj)
	if err != nil {
		return handle(ctx, err)
	}
	return ctx.JSONBlob(http.StatusOK, res)
}

func (controller *UserController) RetrieveUser(ctx echo.Context) error {
	var r requestBody
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err := ctx.Bind(&r); err != nil {
		return handle(ctx, fmt.Errorf("入力をパースできません"))
	}
	if err != nil {
		return handle(ctx, err)
	}
	res, err := controller.interactor.RetrieveUser(id)
	if err != nil {
		return handle(ctx, err)
	}
	return ctx.JSONBlob(http.StatusOK, res)
}

func (controller *UserController) UpdateUser(ctx echo.Context) error {
	var r requestBody
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err := ctx.Bind(&r); err != nil {
		return handle(ctx, fmt.Errorf("入力をパースできません"))
	}
	obj := r.convert()
	if err != nil {
		return handle(ctx, err)
	}
	res, err := controller.interactor.UpdateUser(id, obj)
	if err != nil {
		return handle(ctx, err)
	}
	return ctx.JSONBlob(http.StatusOK, res)
}

func (controller *UserController) DeleteUser(ctx echo.Context) error {
	var r requestBody
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err := ctx.Bind(&r); err != nil {
		return handle(ctx, fmt.Errorf("入力をパースできません"))
	}
	if err != nil {
		return handle(ctx, err)
	}
	res, err := controller.interactor.DeleteUser(id)
	if err != nil {
		return handle(ctx, err)
	}
	return ctx.JSONBlob(http.StatusOK, res)
}
