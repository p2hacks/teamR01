package model

type Rate struct {
	ID			string `json:id`
	RATE		int `json:rate`
}

type SetRate struct {
	STATUS	bool `json:status`
}