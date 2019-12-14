package model

type Present struct {
	ID			string `json:id`
	ASIN		string `json:rate`
}

type User struct {
	ID			string
	PHONE		string `gorm:"type:varchar(11);"`
}

type SetPresent struct {
	STATUS	bool `json:status`
}