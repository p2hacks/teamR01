package model

type Message struct {
	ID      string 
	MESSAGE string  
}

type Pair struct{
	SANTA string
	CHILD string
}

type Request struct{
	ID string
}

type Response struct{
	MESSAGE string
	STATUS string
}