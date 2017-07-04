package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	// var movies []Movie
	// csvFile, _ := os.Open("../../movie_metadata.csv")
	// reader := csv.NewReader(bufio.NewReader(csvFile))
	// for {
	// 	// fmt.Println("hellosrse")
	// 	line, error := reader.Read()
	// 	if error == io.EOF {
	// 		break
	// 	} else if error != nil {
	// 		// fmt.Println("hellofff")
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
