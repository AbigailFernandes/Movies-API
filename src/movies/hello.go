package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	MovieTitle   string `json:"movietitle"`
	DirectorName string `json:"directorname"`
	IMDBScore    string `json:"imdbscore"`
}

type Adult struct {
	Age           string `json:"age"`
	Workclass     string `json:"workclass"`
	Education     string `json:"education"`
	MaritalStatus string `json:"maritalstatus"`
	Occupation    string `json:"occupation"`
	Sex           string `json:"sex"`
	HoursPerWeek  string `json:"hoursperweek"`
	NativeCountry string `json:"nativecountry"`
}

func main() {
	readMovies()
	readAdultData()
}

func readAdultData() {
	// 	age: continuous.
	// workclass: Private, Self-emp-not-inc, Self-emp-inc, Federal-gov, Local-gov, State-gov, Without-pay, Never-worked.
	// fnlwgt: continuous.
	// education: Bachelors, Some-college, 11th, HS-grad, Prof-school, Assoc-acdm, Assoc-voc, 9th, 7th-8th, 12th, Masters, 1st-4th, 10th, Doctorate, 5th-6th, Preschool.
	// education-num: continuous.
	// marital-status: Married-civ-spouse, Divorced, Never-married, Separated, Widowed, Married-spouse-absent, Married-AF-spouse.
	// occupation: Tech-support, Craft-repair, Other-service, Sales, Exec-managerial, Prof-specialty, Handlers-cleaners, Machine-op-inspct, Adm-clerical, Farming-fishing, Transport-moving, Priv-house-serv, Protective-serv, Armed-Forces.
	// relationship: Wife, Own-child, Husband, Not-in-family, Other-relative, Unmarried.
	// race: White, Asian-Pac-Islander, Amer-Indian-Eskimo, Other, Black.
	// sex: Female, Male.
	// capital-gain: continuous.
	// capital-loss: continuous.
	// hours-per-week: continuous.
	// native-country: United-States, Cambodia, England,
	var adults []Adult
	csvFile, _ := os.Open("../adult.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		adults = append(adults, Adult{
			Age:           line[0],
			Workclass:     line[1],
			Education:     line[3],
			MaritalStatus: line[5],
			Occupation:    line[6],
			Sex:           line[9],
			HoursPerWeek:  line[12],
			NativeCountry: line[13],
		})
	}
	// adultJSON, _ := json.Marshal(adults)
	// fmt.Println(len(adults))
	c := make(chan []Adult)
	go findIndiansWithMasters(adults[:len(adults)/2], c)
	go findIndiansWithMasters(adults[len(adults)/2:], c)

	x, y := <-c, <-c
	x = append(x, y...)
	for _, v := range x {
		fmt.Println(v)
	}
}

func findIndiansWithMasters(s []Adult, c chan []Adult) {
	var filteredAdults []Adult
	for _, v := range s {
		if strings.Contains(v.Education, "Masters") && strings.Contains(v.NativeCountry, "India") && strings.Contains(v.Occupation, "Tech-support") {
			filteredAdults = append(filteredAdults, v)
		}
	}
	c <- filteredAdults
}

func readMovies() {
	var movies []Movie
	csvFile, _ := os.Open("../movie_metadata.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		movies = append(movies, Movie{
			MovieTitle:   line[11],
			DirectorName: line[1],
			IMDBScore:    line[25],
		})
	}
	c := make(chan []Movie)
	go findMovies(movies[1:len(movies)/2], 9, c)
	go findMovies(movies[len(movies)/2:], 9, c)
	x, y := <-c, <-c
	x = append(x, y...)
	for _, v := range x {
		fmt.Println(v.MovieTitle)
	}
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
