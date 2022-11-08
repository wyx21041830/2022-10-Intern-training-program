package middleware

import (
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		User.Name = "author"
		User.password = ""
		User.UserId = 0
		return next(c)
	}
}
