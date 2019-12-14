package model

type Message struct {
	ID      string
	MESSAGE string
}

type Pair struct {
	SANTA string `json:"santa"`
	CHILD string `json:"child"`
}

type Request struct {
	ID string
}

type Response struct {
	STATUS  bool
	MESSAGE string
}
