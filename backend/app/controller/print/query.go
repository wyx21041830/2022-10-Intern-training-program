package print

import (
	"backend/App/Response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DelGet(c echo.Context) error {
	name := c.QueryParam("name")
	fmt.Println(name)
	return Response.SendResponse(c, http.StatusOK, "query is ok", "123")
}
