# Movies-API
A REST API created in Golang

# Usage

You should have GoLang installed. <br />
Run go get <br />
go install <br />
Run the exe generated. <br />
Now navigate to http://localhost:8080/

# Endpoints

http://localhost:8080/movies <br />
http://localhost:8080/movies/{movieID} <br />
http://localhost:8080/movies?ratinggeq=9 <br />

Query parameter ratinggeq will give you the list of movies with IMDB rating greater than or equal to the value
