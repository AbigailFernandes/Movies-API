package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var movies []Movie
var currentID int

func init() {
	csvFile, _ := os.Open("D:/ABIGAIL CIMPRESS/Movies-API/Movies-API/src/movie_metadata.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		currentID++
		movies = append(movies, Movie{
			ID:           currentID,
			MovieTitle:   line[11],
			DirectorName: line[1],
			IMDBScore:    line[25],
		})
	}
}

func getMovies(rating int) []Movie {
	if rating != 0 {
		c := make(chan []Movie)
		go findMovies(movies[1:len(movies)/2], 9, c)
		go findMovies(movies[len(movies)/2:], 9, c)
		x, y := <-c, <-c
		x = append(x, y...)
		return x
	}
	return movies
}

func findMovies(s []Movie, rating float64, c chan []Movie) {
	var filteredMovies []Movie
	for _, v := range s {
		score, err := strconv.ParseFloat(v.IMDBScore, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if score > rating {
			filteredMovies = append(filteredMovies, v)
		}
	}
	c <- filteredMovies
}

func getMoviesById(id int) Movie {
	for _, t := range movies {
		if t.ID == id {
			return t
		}
	}
	return Movie{}
}
