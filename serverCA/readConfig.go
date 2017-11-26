package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config reads the config
type Config struct {
	IP             string `json:"ip"`
	Port           string `json:"port"`
	ServerIP       string `json:"serverip"`
	ServerPort     string `json:"serverport"`
	ServerRESTPort string `json:"serverrestport"`
	WebServerPort  string `json:"webserverport"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
