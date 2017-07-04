package main

// type Todo struct {
// 	Name      string    `json:"name"`
// 	Completed bool      `json:"completed"`
// 	Due       time.Time `json:"due"`
// }

type Movie struct {
	MovieTitle   string `json:"movietitle"`
	DirectorName string `json:"directorname"`
	IMDBScore    string `json:"imdbscore"`
}

type Movies []Movie

// type Todos []Todo
