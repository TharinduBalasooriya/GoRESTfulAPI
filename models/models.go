package models

type Movie struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Category   string `json:"category"`
	Year       string `json:"year"`
	ImdbRating string `json:"imdbRating"`
}

type Message struct {
	Message string `json:"message"`
}
