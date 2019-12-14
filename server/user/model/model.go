package model

type Request struct {
	PHONE		string `json:phone` 
}

type User struct {
	ID			string
	PHONE		string `gorm:"type:varchar(11);"`
}

type SetUser struct {
	STATUS	bool `json:status`
	ID		string `json:id`
}