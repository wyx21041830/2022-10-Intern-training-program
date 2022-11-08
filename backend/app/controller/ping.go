package controller

import (
	"flybitch/app/response"
	"net/http"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return response.SendResponse(c, http.StatusOK, "ok", "pong!")
}
