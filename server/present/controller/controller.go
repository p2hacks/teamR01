package controller

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/p2hacks/teamR01/server/present/model"
)

type IsController struct {
	DB	*gorm.DB
}

func (ctrl *IsController)SetPresentInfo(context *gin.Context) {
	var request, present model.Present
	var user model.User

	// var id string
	err := context.BindJSON(&request)
	if err != nil {
		setPresent := model.SetPresent{
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setPresent)
		return
	}

	ctrl.DB.Table("users").Find(&user, "id=?", request.ID)
	if user.ID == "" {
		setPresent := model.SetPresent{
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setPresent)
		return
	}

	present = model.Present{
		ID: user.ID,
		ASIN: request.ASIN,
	}

	ctrl.DB.Table("presents").Create(&present)

	setPresent := model.SetPresent{
		STATUS: true,
	}
	context.JSON(http.StatusOK, setPresent)
}