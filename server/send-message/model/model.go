package model

type Message struct {
	//gorm.Model
	//Model specify the model
	//you would like to run db operions
	ID      string `json:"id"`
	MESSAGE string `json:"message"`
}

type Status struct {
	STATUS bool `json:status`
}
