package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
)

type Peer struct {
	ID       string   `json:"id"` //in the future, this will be the peer hash
	IP       string   `json:"ip"`
	Port     string   `json:"port"`
	RESTPort string   `json:"restport"`
	Role     string   `json:"role"` //client or server
	Conn     net.Conn `json:"conn"`
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

	//read the stored blockchain
	err := blockchain.readFromDisk()
	check(err)
	blockchain.print()

	//runningPeer.ID = strconv.Itoa(randInt(1, 1000)) //0 is reserved for server
	runningPeer.IP = config.IP
	runningPeer.Port = config.Port
	runningPeer.RESTPort = config.RESTPort
	runningPeer.ID = hashPeer(runningPeer)
	runningPeer.Role = "client"

	//TODO clean and reorder the following lines (43 to 62)
	//read flags, to know if is runned as p2p server
	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			color.Yellow("Running as p2p server")
			runningPeer.Role = "server"
			runningPeer.Port = config.ServerPort
			runningPeer.ID = hashPeer(runningPeer)
			//runningPeer.ID = "0"
		}
		if len(os.Args) > 3 {
			config.Port = os.Args[2]
			config.RESTPort = os.Args[3]
			runningPeer.Port = os.Args[2]
			runningPeer.RESTPort = os.Args[3]
			if os.Args[1] == "server" {
				runningPeer.Port = config.ServerPort
			}
			runningPeer.ID = hashPeer(runningPeer)
		}
	}

	go runRestServer()
	thisPeerID = runningPeer.ID
	outcomingPeersList.PeerID = runningPeer.ID
	fmt.Println(runningPeer)
	//outcomingPeersList.Peers = append(outcomingPeersList.Peers, runningPeer)
	outcomingPeersList = appendPeerIfNoExist(outcomingPeersList, runningPeer)
	fmt.Println(outcomingPeersList)
	if runningPeer.Role == "server" {
		go acceptPeers(runningPeer)
	}
	if runningPeer.Role == "client" {
		var serverPeer Peer
		//serverPeer.ID = "0"
		serverPeer.IP = config.ServerIP
		serverPeer.Port = config.ServerPort
		serverPeer.ID = hashPeer(serverPeer)
		serverPeer.Role = "server"
		go acceptPeers(runningPeer)
		connectToPeer(serverPeer)
		reconstructBlockchainFromBlock("http://"+config.IP+":"+config.ServerRESTPort, blockchain.GenesisBlock)
	}

	for running {
		time.Sleep(1000 * time.Millisecond)
	}
}
