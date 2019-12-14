package controller

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/p2hacks/teamR01/server/rate/model"
)

type IsController struct {
	DB	*gorm.DB
}

// test struct
type User struct {
	ID			string
	PHONE		string `gorm:"type:varchar(11);"`
}

func (ctrl *IsController)SetRateInfo(context *gin.Context) {
	var request, rate model.Rate
	var user User

	// var id string
	err := context.BindJSON(&request)
	if err != nil {
		setRate := model.SetRate{
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setRate)
		return
	}

	ctrl.DB.Table("users").Find(&user, "id=?", request.ID)
	if user.ID == "" {
		setRate := model.SetRate{
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setRate)
		return
	}

	rate = model.Rate{
		ID: user.ID,
		RATE: request.RATE,
	}

	ctrl.DB.Table("rates").Create(&rate)

	setRate := model.SetRate{
		STATUS: true,
	}
	context.JSON(http.StatusOK, setRate)
}