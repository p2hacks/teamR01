package controller

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/p2hacks/teamR01/server/send-message/model"
	"net/http"
)

type IsController struct {
	DB *gorm.DB
}

/*TODO:
POSTでユーザーIDと送るメッセージを受け取る
結果をJSONでstatusを返す。
*/

func (ctrl *IsController) MessageHandler(c echo.Context) error {
	var post model.Message
	var res model.Status
	err := c.Bind(&post)
	if err != nil {
		res.STATUS = false
		return c.JSON(http.StatusBadRequest, res)
	}
	ctrl.DB.Create(&post)
	res.STATUS = true
	return c.JSON(http.StatusOK, res)
}

func InitController(db *gorm.DB) IsController {
	return IsController{DB: db}
}
