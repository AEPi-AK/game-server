package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	var routes = Routes{
		Route{
			"Poll",
			"POST",
			"/poll",
			Poll,
		},
		Route{
			"Hello",
			"POST",
			"/hello",
			Hello,
		},
		Route{
			"Attack",
			"POST",
			"/attack",
			Attack,
		},
		Route{
			"HelloMonster",
			"POST",
			"/hello-monster",
			HelloMonster,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
