package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MovieShow",
		"GET",
		"/movies",
		MovieShow,
	},
	Route{
		"MovieIndex",
		"GET",
		"/movies/{movieID}",
		MovieIndex,
	},
}
