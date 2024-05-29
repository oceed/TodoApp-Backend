package models

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
