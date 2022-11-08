package controller

import (
	"backend/App/Response"
	"backend/model"
	"backend/services"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// 创建
func Create(c echo.Context) error {
	todo := new(model.Todo)
	defer c.Request().Body.Close()
	id, err := io.ReadAll(c.Request().Body)
	if err == nil {
		logrus.Error(err)
	}
	if err := json.Unmarshal(id, &todo); err != nil {
		logrus.Error(err)
	}
	services.TodoAdd(todo)
	return Response.SendRes(c, http.StatusOK, "post is ok", "add is ok")
}

// Update
func Update(c echo.Context) error {
	todo := new(model.Todo)
	defer c.Request().Body.Close()
	id, err := io.ReadAll(c.Request().Body)
	if err == nil {
		logrus.Error("error1")
	}
	if err := json.Unmarshal(id, &todo); err != nil {
		logrus.Error("error2")
	}
	services.TodoUpdate(todo)
	return Response.SendRes(c, http.StatusOK, "post is ok", "update is ok")
}

// Select
func Select(c echo.Context) error {
	id := c.QueryParam("id")
	i, _ := strconv.Atoi(id)
	num := int(i)
	tmp := services.TodoSelect(num)
	return Response.SendRes(c, http.StatusOK, "post is ok", tmp)
} // Delete
func Del(c echo.Context) error {
	id := c.QueryParam("id")
	i, _ := strconv.Atoi(id)
	num := int(i)
	services.TodoDelete(num)
	return Response.SendRes(c, http.StatusOK, "post is ok", "delete is ok")
}
