package main

import (
	"encoding/json"
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

type Address struct {
	Address string `json:"address"` //the pubK of the user, to perform logins
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CA")
}
func GetPeers(w http.ResponseWriter, r *http.Request) {
	getPeers()

	jResp, err := json.Marshal(peersList)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	fmt.Print("aaaaa: ")
	fmt.Println(blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	reconstructBlockchainFromBlock("http://"+config.IP+":"+config.ServerRESTPort, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	jResp, err := json.Marshal(blockchain)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
