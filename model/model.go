package model

type Articulo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Date string `json:"date"`
}


type BodyPostRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Date string `json:"date"`
}