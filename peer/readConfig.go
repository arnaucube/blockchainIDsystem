package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config reads the config
type Config struct {
	IP         string `json:"ip"`
	Port       string `json:"port"`
	RestIP     string `json:"restip"`
	RestPort   string `json:"restport"`
	ServerIP   string `json:"serverip"`
	ServerPort string `json:"serverport"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
