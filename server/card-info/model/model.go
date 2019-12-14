package model

type Card struct {
	ID			string
	NUM			string
	YEAR		string
	MONTH		string
	CODE		string
	NAME		string
}

type User struct {
	ID			string
	PHONE		string
}

type SetCardInfo struct {
	STATUS		bool `json:status`
}