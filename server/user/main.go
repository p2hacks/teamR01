package main

import(
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/p2hacks/teamR01/server/user/config"
	"github.com/p2hacks/teamR01/server/user/controller"
	"github.com/p2hacks/teamR01/server/user/model"
)

func main() {
	var ctrl controller.IsController
	db, err := initializeDataBase()
	if err != nil {
		log.Fatalf("failed initialize db. err=%s", err)
	}

	ctrl = initializeController(db)
	router := setupRouter(ctrl)
	err = router.Run(":9000")
	if err != nil {
		log.Fatalf("failed launch router. err=%s", err)
	}
}

func initializeDataBase() (*gorm.DB, error){
	var db *gorm.DB
	var err error
	var count time.Duration
	token := config.GetConnectionToken()

	count = 1
	for {
		if count > 5 {
			return nil, fmt.Errorf("faild mysql connection")
		}
		db, err = gorm.Open("mysql", token)
		if err == nil {
			db.AutoMigrate(&model.User{})
			return db, nil
		}
		time.Sleep(3 * time.Second)

		count++
	}

	return nil, err
}

func initializeController(db *gorm.DB) (controller.IsController){
	return controller.IsController{DB: db}
}

func setupRouter(ctrl controller.IsController) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "This is User API"})
	})
	router.POST("/set", ctrl.SetUserInfo)
	return router
}