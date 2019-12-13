package controller

import(
	"fmt"
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
	var user, test model.User
	var id []byte

	// var id string
	err := context.BindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"err": fmt.Sprintf("failed to unmarshal json data. err=%s", err),
		})
		return
	}

	id, err = service.CreateUserId(request.PHONE)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"err": fmt.Sprintf("failed to unmarshal json data. err=%s", err),
		})
		return
	}

	user = model.User{
		ID: id,
		PHONE: request.PHONE,
	}

	ctrl.DB.Create(&user)

	ctrl.DB.Find(&test, "id=?", id)

	context.JSON(http.StatusOK, gin.H{"status": true, "err": test})
}