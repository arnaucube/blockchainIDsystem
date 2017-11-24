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
	ID   string   `json:"id"` //in the future, this will be the peer hash
	IP   string   `json:"ip"`
	Port string   `json:"port"`
	Role string   `json:"role"` //client or server
	Conn net.Conn `json:"conn"`
}

var running bool
var thisPeerID string
var runningPeer Peer

func main() {
	//initialize some vars
	rand.Seed(time.Now().Unix())
	running = true

	color.Blue("Starting Peer")
	//read configuration file
	readConfig("config.json")

	runningPeer.ID = strconv.Itoa(randInt(1, 1000)) //0 is reserved for server
	runningPeer.IP = config.IP
	runningPeer.Port = config.Port
	runningPeer.Role = "client"

	go runRestServer()

	//read flags, to know if is runned as p2p server
	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			color.Yellow("Running as p2p server")
			runningPeer.Role = "server"
			runningPeer.Port = config.ServerPort
			runningPeer.ID = "0"
		}
	}
	thisPeerID = runningPeer.ID
	peersList.PeerID = runningPeer.ID
	fmt.Println(runningPeer)
	peersList.Peers = append(peersList.Peers, runningPeer)
	fmt.Println(peersList)
	if runningPeer.Role == "server" {
		go acceptPeers(runningPeer)
	}
	if runningPeer.Role == "client" {
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
