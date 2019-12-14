package model

type Message struct {
	//gorm.Model
	//Model specify the model
	//you would like to run db operions
	ID      string
	MESSAGE string
}

type Status struct {
	STATUS bool `json:status`
}
