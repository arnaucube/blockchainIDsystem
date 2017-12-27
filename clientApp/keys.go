package main

import (
	"encoding/json"
	"io/ioutil"

	ownrsa "./ownrsa"
)

func readKeys(path string) []ownrsa.PackRSA {
	var keys []ownrsa.PackRSA

	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &keys)

	return keys
}

func saveKeys(keys []ownrsa.PackRSA, path string) {
	jsonKeys, err := json.Marshal(keys)
	check(err)
	err = ioutil.WriteFile(path, jsonKeys, 0644)
	check(err)
}
