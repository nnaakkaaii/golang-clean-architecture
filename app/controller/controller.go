package controller

import (
	"clean_architecture/app/interactor"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	interactor *interactor.UserInteractorInterface
	presenter  *UserPresenterInterface
}

func NewUserController(interactor *interactor.UserInteractorInterface, presenter *UserPresenterInterface) *UserController {
	return &UserController{
		interactor: interactor,
		presenter:  presenter,
	}
}

var _ UserControllerInterface = (*UserController)(nil)

func (controller *UserController) ListUsers(ctx echo.Context) error {
	users, err := (*controller.interactor).ListUsers()
	if err != nil {
		return handle(ctx, err)
	}
	res, err := (*controller.presenter).ListUsersView(users)
	if err != nil {
		return handle(ctx, err)
	}
	return ctx.JSONBlob(http.StatusOK, res)
}

type Date struct {
	string
}

func (d *Date) Time() (date time.Time, err error) {
	date, err = time.Parse("2006-01-02", d.string)
	return
}

type requestBody struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	BirthDate string `json:"birth_date" form:"birth_date"`
}

func (r *requestBody) convert() (req *interactor.UserRequest, err error) {
	date := &Date{r.BirthDate}
	t, err := date.Time()
	if err != nil {
		return
	}
	req = &interactor.UserRequest{
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		BirthYear:  t.Year(),
		BirthMonth: int(t.Month()),
		BirthDay:   t.Day(),
	}
	return
}

func (controller *UserController) CreateUser(ctx echo.Context) error {
	var r requestBody
	if err := ctx.Bind(&r); err != nil {
		return handle(ctx, err)
	}
	obj, err := r.convert()
	if err != nil {
		return handle(ctx, err)
	}
	user, err := (*controller.interactor).CreateUser(obj)
	if err != nil {
		return handle(ctx, err)
	}
	res, err := (*controller.presenter).CreateUserView(user)
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
	user, err := (*controller.interactor).RetrieveUser(id)
	if err != nil {
		return handle(ctx, err)
	}
	res, err := (*controller.presenter).RetrieveUserView(user)
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
	obj, err := r.convert()
	if err != nil {
		return handle(ctx, err)
	}
	user, err := (*controller.interactor).UpdateUser(id, obj)
	if err != nil {
		return handle(ctx, err)
	}
	res, err := (*controller.presenter).UpdateUserView(user)
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
	user, err := (*controller.interactor).DeleteUser(id)
	if err != nil {
		return handle(ctx, err)
	}
	res, err := (*controller.presenter).DeleteUserView(user)
	if err != nil {
		return handle(ctx, err)
	}
	return ctx.JSONBlob(http.StatusOK, res)
}
