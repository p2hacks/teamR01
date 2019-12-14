package dbconnect

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/p2hacks/teamR01/server/send-message/config"
	"github.com/p2hacks/teamR01/server/send-message/model"
)

func InitDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var count time.Duration
	token := config.GetConnectionToken()

	count = 1
	for {
		if count > 5 {
			return nil, fmt.Errorf("データベース接続に失敗しました！")
		}
		db, err = gorm.Open("mysql", token)
		if err == nil {
			db.AutoMigrate(&model.Message{})
			return db, nil
		}
		time.Sleep(3 * time.Second)

		count++
	}
	db.LogMode(true)
	return nil, err
}

