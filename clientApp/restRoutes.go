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
		"IDs",
		"GET",
		"/ids",
		IDs,
	},
	Route{
		"NewID",
		"GET",
		"/newid",
		NewID,
	},
	Route{
		"BlindAndVerify",
		"GET",
		"/blindandverify/{pubK}",
		BlindAndVerify,
	},
}
