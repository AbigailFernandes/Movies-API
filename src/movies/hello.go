package main

import (
	"log"
	"net/http"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// type Movie struct {
// 	MovieTitle   string `json:"movietitle"`
// 	DirectorName string `json:"directorname"`
// 	IMDBScore    string `json:"imdbscore"`
// }

func main() {
	// router := NewRouter()
	// log.Fatal(http.ListenAndServe(":8080", router))
	// var movies []Movie
	// csvFile, _ := os.Open("../movie_metadata.csv")
	// reader := csv.NewReader(bufio.NewReader(csvFile))
	// for {
	// 	line, error := reader.Read()
	// 	if error == io.EOF {
	// 		break
	// 	} else if error != nil {
	// 		log.Fatal(error)
	// 	}
	// 	movies = append(movies, Movie{
	// 		MovieTitle:   line[11],
	// 		DirectorName: line[1],
	// 		IMDBScore:    line[25],
	// 	})
	// }
	// movieJson, _ := json.Marshal(movies)
	// fmt.Println(string(movieJson))
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, string(movieJson))
	// })
	// log.Fatal(http.ListenAndServe(":8080", router))
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

// func Index(w http.ResponseWriter, r *http.Request) {
//
// fmt.Println(string(movieJson))
// fmt.Fprintf(w, string(movieJson))
// }
