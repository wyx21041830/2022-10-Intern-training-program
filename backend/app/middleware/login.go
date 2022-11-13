package middleware

//中间件
import (
	"backend/model"

	"github.com/labstack/echo/v4"
)

var User model.Users

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		User.Name = "author"
		User.Password = ""
		User.ID = 0
		return next(c)
	}
}
