package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
)

type Peer struct {
	IP   string
	Port string
	Role string //client or server
	Conn net.Conn
}
type Msg struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

var listPeers []Peer
var running bool

func main() {
	//initialize some vars
	running = true
	var peer Peer

	color.Blue("Starting Peer")
	readConfig("config.json")

	peer.IP = config.IP
	peer.Port = config.Port
	peer.Role = "client"

	//read flags, to know if is runned as p2p server
	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			color.Yellow("Running as p2p server")
			peer.Role = "server"
			peer.Port = config.ServerPort
			go acceptPeers(peer)
		}
	}
	fmt.Println(peer)
	if peer.Role == "client" {
		var newPeer Peer
		newPeer.IP = config.ServerIP
		newPeer.Port = config.ServerPort
		newPeer.Role = "server"
		connectToPeer(newPeer)
	}

	for running {
		time.Sleep(1000 * time.Millisecond)
	}
}
