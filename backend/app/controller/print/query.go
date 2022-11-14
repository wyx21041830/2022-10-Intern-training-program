package print

import (
	"backend/App/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DelGet(c echo.Context) error {
	name := c.QueryParam("name")
	fmt.Println(name)
	return response.SendResponse(c, http.StatusOK, "query", "1")
}
