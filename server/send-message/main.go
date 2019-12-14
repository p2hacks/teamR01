package main

import(
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/p2hacks/teamR01/server/send-message/controller"
	"github.com/p2hacks/teamR01/server/send-message/dbconnect"
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
        return c.String(http.StatusOK, "This is Send-message API")
    })
	e.POST("/set",ctrl.MessageHandler)//クライアントからuser_id,messageを受け取るルータ
	//TODO サーバーを立ち上げる
	e.Logger.Fatal(e.Start(":9003"))
}