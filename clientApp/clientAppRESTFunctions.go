package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	ownrsa "./ownrsa"
)

//generate key pair
//blind m
//unblind m

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "serverIDsigner")
}

func IDs(w http.ResponseWriter, r *http.Request) {
	//read the keys stored in /keys directory
	keys := readKeys("keys.json")
	saveKeys(keys, "keys.json")

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
func NewID(w http.ResponseWriter, r *http.Request) {
	//generate RSA keys pair
	newKey := ownrsa.GenerateKeyPair()

	key := ownrsa.PackKey(newKey)
	fmt.Println(key)

	keys := readKeys("keys.json")
	keys = append(keys, key)
	saveKeys(keys, "keys.json")

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
