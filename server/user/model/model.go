package model

type Request struct {
	NAME		string `json:name`
	POSTALF		string `json:postalf`
	POSTALL		string `json:postall`
	STATE		string `json:state`
	ADDRESS		string `json:address`
	PHONE		string `json:phone`	
}

type User struct {
	ID			string `gorm:"primary_key"`
	NAME		string `gorm:"type:varchar(20);"`
	POSTALF		string `gorm:"type:varchar(3);"`
	POSTALL		string `gorm:"type:varchar(4);"`
	STATE		string `gorm:"type:varchar(3);"`
	ADDRESS		string `gorm:"type:varchar(256);"`
	PHONE		string `gorm:"type:varchar(20);"`
}