package model

type Message struct {
	//gorm.Model
	//Model specify the model
	//you would like to run db operions
	ID      string `json:"user_id" gorm:"primary_key" "column:"user_id`
	MESSAGE string `json:"message" gorm:"column:message"`
}

type Status struct {
	STATUS bool `json:status`
}
