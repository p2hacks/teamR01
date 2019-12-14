package controller

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/p2hacks/teamR01/server/pair/model"
)

type IsController struct {
	DB	*gorm.DB
}

func (ctrl *IsController)GetPairInfo(context *gin.Context) {
	var request model.User
	var userRate, targetRate model.Rate
	var rates []model.Rate
	var present model.Present

	// var id string
	err := context.BindJSON(&request)
	if err != nil {
		setPair := model.SetPair{
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setPair)
		return
	}

	ctrl.DB.Table("rates").Find(&userRate, "id=?", request.ID)
	if userRate.ID == "" {
		setPair := model.SetPair{
			STATUS: false, 
		}
		context.JSON(http.StatusInternalServerError, setPair)
		return
	}
	ctrl.DB.Table("rates").Not("id", userRate.ID).Find(&rates, "rate=?", userRate.RATE)
	if userRate.ID == "" {
		setPair := model.SetPair{
			STATUS: false,
		}
		context.JSON(http.StatusInternalServerError, setPair)
		return
	}
	targetRate.ID = rates[0].ID
	targetRate.RATE = rates[0].RATE

	pair := model.Pair{
		SANTA: userRate.ID,
		CHILD: targetRate.ID,
	}

	ctrl.DB.Table("rates").Delete(&targetRate)	
	ctrl.DB.Table("pairs").Create(&pair)

	ctrl.DB.Table("presents").Find(&present, "id=?", pair.CHILD)

	setPair := model.SetPair{
		STATUS: true,
		ASIN: present.ASIN,
	}
	context.JSON(http.StatusOK, setPair)
}