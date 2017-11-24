package main

import (
	"fmt"
	"net/http"
)

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	//ipFilter(w, r)
	fmt.Fprintln(w, runningPeer.ID)
}
