package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
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
		"PostUser",
		"POST",
		"/register",
		PostUser,
	},
}

type Address struct {
	Address string `json:"address"` //the pubK of the user, to perform logins
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, runningPeer.ID)
}
func GetPeers(w http.ResponseWriter, r *http.Request) {
	jResp, err := json.Marshal(outcomingPeersList)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
func PostUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var address Address
	err := decoder.Decode(&address)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Println(address)
	color.Blue(address.Address)

	//TODO add the verification of the address, to decide if it's accepted to create a new Block
	block := blockchain.createBlock(address)
	blockchain.addBlock(block)

	go propagateBlock(block)

	jResp, err := json.Marshal(blockchain)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
