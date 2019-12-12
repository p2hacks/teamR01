package main

import(
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/p2hacks/teamR01/server/send-message/controller"
	"github.com/p2hacks/teamR01/server/send-message/model"
	"github.com/p2hacks/teamR01/server/send-message/config"
	"github.com/p2hacks/teamR01/server/send-message/dbconnect"
)

func main(){
	var ctrl controller.IsController
	//creating new instance
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//TODO データベースの初期化
	db, err :=connect.InitDB()
	if err != nil{
	}
	ctrl = controller.InitController(db)
	//TODO ルーティング
	//e.GET("/",hogehoge)//bff用
	e.POST("/send",ctrl.controller.MessageHandler)//クライアントからuser_id,messageを受け取るルータ
	//TODO サーバーを立ち上げる
	e.Logger.Fatal(e.Start(":8080"))
}