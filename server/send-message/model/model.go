package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID      string `json:"user_id"`
	MESSAGE string `json:"message"`
}

type Status struct {
	STATUS bool `json:status`
}
