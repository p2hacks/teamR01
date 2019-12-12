package main

import (
	"github.com/labstack/echo"
	github.com/labstack/echo/middleware"
)

func main() {
	//creating new instance
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//TODO データベースの初期化

	//TODO ルーティング
	e.GET("/")
	e.POST("/send")
	//TODO サーバーを立ち上げる
	e.Logger.Fatal(e.Start(":8080"))
}
