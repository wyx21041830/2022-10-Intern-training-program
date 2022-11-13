package app

import (
	"backend/app/controller"
	"backend/app/middleware"
)

func addRoutes() {
	api := e.Group("api")
	api.Use(middleware.Auth)
	api.GET("/ping", controller.Ping)
	api.POST("/create", controller.Create)
	api.POST("/update", controller.Update)
	api.GET("/delete", controller.Del)
	api.GET("/select", controller.Select)
}
