package model

type User struct {
	ID			string `json:id`
}

type Rate struct {
	ID			string
	RATE		int
}

type Pair struct {
	SANTA		string
	CHILD		string
}

type Present struct {
	ID			string
	ASIN		string
}

type SetPair struct {
	STATUS	bool `json:status`
	ASIN	string `json:asin`
}