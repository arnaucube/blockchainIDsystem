package main

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetPeers",
		"GET",
		"/peers",
		GetPeers,
	},
	Route{
		"GetBlockchain",
		"GET",
		"/blockchain",
		GetBlockchain,
	},
}
