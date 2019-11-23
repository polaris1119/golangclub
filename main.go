package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/polaris1119/golangclub/http/controller"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// 去除尾部斜杠
	e.Pre(middleware.RemoveTrailingSlash())

	// 服务静态文件
	e.Static("/static", "static")

	controller.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":2019"))
}
