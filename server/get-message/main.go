package main

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/p2hacks/teamR01/server/get-message/controller"
	"github.com/p2hacks/teamR01/server/get-message/dbconnect"
	"net/http"
)

func main(){
	var ctrl controller.IsController
	//creating new instance
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//TODO データベースの初期化
	db, err := dbconnect.InitDB()
	if err != nil{
	}
	ctrl = controller.InitController(db)
	//TODO ルーティング
	e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "This is Get-Message API")
    })
	e.POST("/message",ctrl.MessageHandler)
	//TODO サーバーを立ち上げる
	e.Logger.Fatal(e.Start(":8080"))
}