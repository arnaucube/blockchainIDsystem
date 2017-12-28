package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config reads the config
type Config struct {
	Port           string `json:"port"`
	KeysDirectory  string `json:"keysDirectory"`
	ServerIDSigner Server `json:"serverIDsigner"`
}
type Server struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
