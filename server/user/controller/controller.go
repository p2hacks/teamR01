package controller

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/p2hacks/teamR01/server/user/model"
	"github.com/p2hacks/teamR01/server/user/service"
)

type IsController struct {
	DB	*gorm.DB
}

func (ctrl *IsController)SetUserInfo(context *gin.Context) {
	var request model.Request
	var user model.User
	var id string

	// var id string
	err := context.BindJSON(&request)
	if err != nil {
		setUser := model.SetUser{
			STATUS: false,
			ID: "",
		}
		context.JSON(http.StatusInternalServerError, gin.H{"setUser": setUser})
		return
	}

	id, err = service.CreateUserId(request.PHONE)
	if err != nil {
		setUser := model.SetUser{
			STATUS: false,
			ID: "",
		}
		context.JSON(http.StatusInternalServerError, gin.H{"setUser": setUser})
		return
	}

	user = model.User{
		ID: id,
		PHONE: request.PHONE,
	}

	ctrl.DB.Create(&user)

	setUser := model.SetUser{
		STATUS: true,
		ID: id,
	}
	context.JSON(http.StatusOK, setUser)
}