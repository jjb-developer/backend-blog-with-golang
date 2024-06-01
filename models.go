package main

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Date string `json:"date"`
}

type BodyRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Date string `json:"date"`
}