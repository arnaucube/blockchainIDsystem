package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Peer struct {
	ID   string   `json:"id"`
	IP   string   `json:"ip"`
	Port string   `json:"port"`
	Role string   `json:"role"` //client or server
	Conn net.Conn `json:"conn"`
}

var running bool
var thisPeerID string

func main() {
	//initialize some vars
	rand.Seed(time.Now().Unix())
	running = true
	var peer Peer

	color.Blue("Starting Peer")
	readConfig("config.json")

	peer.ID = strconv.Itoa(randInt(1, 1000)) //0 is reserved for server
	peer.IP = config.IP
	peer.Port = config.Port
	peer.Role = "client"

	//read flags, to know if is runned as p2p server
	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			color.Yellow("Running as p2p server")
			peer.Role = "server"
			peer.Port = config.ServerPort
			peer.ID = "0"
		}
	}
	thisPeerID = peer.ID
	peersList.PeerID = peer.ID
	fmt.Println(peer)
	peersList.Peers = append(peersList.Peers, peer)
	fmt.Println(peersList)
	if peer.Role == "server" {
		go acceptPeers(peer)
	}
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
