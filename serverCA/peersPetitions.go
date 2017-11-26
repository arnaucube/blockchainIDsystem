package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getPeers() {
	res, err := http.Get("http://" + config.IP + ":" + config.ServerRESTPort + "/peers")
	check(err)
	body, err := ioutil.ReadAll(res.Body)
	check(err)
	err = json.Unmarshal(body, &peersList)
	check(err)
}
