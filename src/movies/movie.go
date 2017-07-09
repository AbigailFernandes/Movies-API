package main

type Movie struct {
	ID           int    `json:"id"`
	MovieTitle   string `json:"movietitle"`
	DirectorName string `json:"directorname"`
	IMDBScore    string `json:"imdbscore"`
}

type Movies []Movie
