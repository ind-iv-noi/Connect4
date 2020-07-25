package main

type Game struct {
	PlayerTurn int       `json:"PlayerTurn"`
	Table      [6][7]int `json:"Table"`
}

type HTTPResponse struct {
	Code     int      `json:"Code"`
	Response []string `json:"Response"`
}
