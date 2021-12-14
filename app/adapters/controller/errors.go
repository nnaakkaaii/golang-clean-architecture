package controller

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// handle ではusecase/errors.goで発生したerrorが送り込まれるので、ここで分類してresponseの場合分けを行う
func handle(ctx echo.Context, err error) error {
	log.Print("API:", err.Error())
	switch {
	default:
		return ctx.JSON(http.StatusInternalServerError, errResp{"unknown server error"})
	}
}

type errResp struct {
	Message string `json:"message"`
}
