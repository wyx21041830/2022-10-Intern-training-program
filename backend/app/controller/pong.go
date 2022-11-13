package controller

import (
	"backend/app/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return response.SendResponse(c, http.StatusOK, "ok", "pong!")
}
