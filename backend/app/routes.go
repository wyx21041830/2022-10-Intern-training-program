package app

import (
	"backend/app/controller/print"

	"backend/app/controller"
	"backend/app/middleware"
)

// 路由
func addRoutes() {
	api := e.Group("api") //定义路由组
	api.Use(middleware.Auth)
	api.GET("/ping", controller.Pong)
	api.POST("/create", print.Create)
	api.POST("/update", print.Update)
	api.GET("/delete", print.Del)
	api.GET("/select", print.Select)
}
