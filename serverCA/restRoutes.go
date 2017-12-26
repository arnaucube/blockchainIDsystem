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
	/*
		POST	/signup
		POST	/loginuser
		POST	/blindsign
		POST	/verifysign

	*/
	Route{
		"Signup",
		"POST",
		"/signup",
		Signup,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"BlindSign",
		"POST",
		"/blindsign",
		BlindSign,
	},
	Route{
		"VerifySign",
		"POST",
		"/verifysign",
		VerifySign,
	},
}
