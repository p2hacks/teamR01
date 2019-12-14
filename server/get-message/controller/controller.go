package controller

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/p2hacks/teamR01/server/get-message/model"
)

type IsController struct {
	DB *gorm.DB
}

/*TODO:
POSTで受け取ったidでpairsから、マッチングしているidを取得する。
取得したidでmessagesから該当するmessageを取得する
結果をJSONでmessageとstatusを返す。
*/

func (ctrl *IsController) MessageHandler(c echo.Context) error {
	var req model.Request
	var pair model.Pair
	var message model.Message
	var res model.Response

	//TODO: POSTで受け取ったidをreqに入れる
	err := c.Bind(&req)
	if err != nil {
		res.STATUS = false
		return c.JSON(http.StatusBadRequest, res)
	}
	//TODO: reqをSANTAに代入、SANTAでpairからCHILDを引っ張ってくる
	ctrl.DB.Table("pairs").Find(&pair, "santa=?", req.ID)

	//CHILDでmessageを検索
	ctrl.DB.Table("messages").First(&message, "id=?", pair.CHILD)

	//Responseとして返すmessageとstatusを更新する。
	res.MESSAGE = message.MESSAGE
	res.STATUS = true

	return c.JSON(http.StatusOK, res)
}

func InitController(db *gorm.DB) IsController {
	return IsController{DB: db}
}
