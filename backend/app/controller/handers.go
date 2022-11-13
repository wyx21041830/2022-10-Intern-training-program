package controller

import (
	"backend/app/response"
	"backend/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// 创建
func Create(c echo.Context) error {
	todo := new(model.Todos)
	defer c.Request().Body.Close()
	id, err := io.ReadAll(c.Request().Body)
	if err == nil {
		logrus.Error(err)
	}
	if err := json.Unmarshal(id, &todo); err != nil {
		logrus.Error(err)
	}
	model.Add(todo)
	return response.SendResponse(c, http.StatusOK, "post is ok", "add is ok")
}

// Update
func Update(c echo.Context) error {
	todo := new(model.Todos)
	defer c.Request().Body.Close()
	id, err := io.ReadAll(c.Request().Body)
	if err == nil {
		logrus.Error("error1")
	}
	if err := json.Unmarshal(id, &todo); err != nil {
		logrus.Error("error2")
	}
	model.Update(todo)
	return response.SendResponse(c, http.StatusOK, "post is ok", "update is ok")
}

// Select
func Select(c echo.Context) error {
	id := c.QueryParam("id")
	i, _ := strconv.Atoi(id)
	num := uint(i)
	tmp := model.Query1(num)
	return response.SendResponse(c, http.StatusOK, "post is ok", tmp)
} // Delete
func Del(c echo.Context) error {
	id := c.QueryParam("id")
	i, _ := strconv.Atoi(id)
	num := uint(i)
	model.Dels(num)
	return response.SendResponse(c, http.StatusOK, "post is ok", "delete is ok")
}
