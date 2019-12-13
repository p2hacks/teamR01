package model

type Request struct {
	PHONE		string `json:phone`	
}

type User struct {
	ID			[]byte
	PHONE		string `gorm:"type:varchar(11);"`
}