package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"
)

type Peer struct {
	ID       string   `json:"id"` //in the future, this will be the peer hash
	IP       string   `json:"ip"`
	Port     string   `json:"port"`
	RESTPort string   `json:"restport"`
	Role     string   `json:"role"` //client or server
	Conn     net.Conn `json:"conn"`
}
type PeersList struct {
	PeerID string
	Peers  []Peer    `json:"peerslist"`
	Date   time.Time `json:"date"`
}

var peersList PeersList

func main() {
	color.Blue("Starting CA")

	//read configuration file
	readConfig("config.json")

	//run thw webserver
	go webserver()

	//run API
	log.Println("api server running")
	log.Print("port: ")
	log.Println(config.Port)
	router := NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+config.Port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func webserver() {
	log.Println("webserver in port " + config.WebServerPort)
	http.Handle("/", http.FileServer(http.Dir("./webapp")))
	http.ListenAndServe(":"+config.WebServerPort, nil)
}
