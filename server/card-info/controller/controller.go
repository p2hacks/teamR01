package controller

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/p2hacks/teamR01/server/card-info/model"
)

type IsController struct {
	DB	*gorm.DB
}

func (ctrl *IsController)SetCardInfo(context *gin.Context) {
	var request model.Card
	var user model.User

	err := context.BindJSON(&request)
	if err != nil {
		setCardInfo := model.SetCardInfo {
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setCardInfo)
		return
	}

	ctrl.DB.Table("users").Find(&user, "id=?", request.ID)
	if user.ID == "" {
		setCardInfo := model.SetCardInfo {
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setCardInfo)
		return
	}

	ctrl.DB.Table("cards").Create(&request)

	setCardInfo := model.SetCardInfo {
		STATUS: true,
	}
	context.JSON(http.StatusOK, setCardInfo)
}