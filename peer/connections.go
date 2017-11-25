package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
)

func acceptPeers(peer Peer) {
	fmt.Println("accepting peers at: " + peer.Port)
	l, err := net.Listen("tcp", peer.IP+":"+peer.Port)
	if err != nil {
		log.Println("Error accepting peers. Listening port: " + peer.Port)
		running = false
	}
	for running {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting peers. Error accepting connection")
			running = false
		}
		var newPeer Peer
		newPeer.IP = getIPFromConn(conn)
		newPeer.Port = getPortFromConn(conn)
		newPeer.Conn = conn
		/*
			//ask to the peer, for the peer ID
			resp, err := http.Get("http://" + newPeer.IP + ":" + newPeer.Port)
			check(err)
			color.Red("-----")
			fmt.Println(resp)
			//newPeer.ID = resp
		*/
		//incomingPeersList.Peers = append(incomingPeersList.Peers, newPeer)
		incomingPeersList = appendPeerIfNoExist(incomingPeersList, newPeer)
		go handleConn(conn, newPeer)
	}
}
func connectToPeer(peer Peer) {
	color.Green("connecting to new peer")
	log.Println("Connecting to new peer: " + peer.IP + ":" + peer.Port)
	conn, err := net.Dial("tcp", peer.IP+":"+peer.Port)
	if err != nil {
		log.Println("Error connecting to: " + peer.IP + ":" + peer.Port)
		return
	}
	peer.Conn = conn
	//outcomingPeersList.Peers = append(outcomingPeersList.Peers, peer)
	outcomingPeersList = appendPeerIfNoExist(outcomingPeersList, peer)
	go handleConn(conn, peer)
}
func handleConn(conn net.Conn, connPeer Peer) {
	connRunning := true
	log.Println("handling conn: " + conn.RemoteAddr().String())
	//reply to the conn, send the peerList
	var msg Msg
	msg.construct("PeersList", "here my outcomingPeersList")
	msg.PeersList = outcomingPeersList
	msgB := msg.toBytes()
	_, err := conn.Write(msgB)
	check(err)

	for connRunning {
		/*
			buffer := make([]byte, 1024)
			bytesRead, err := conn.Read(buffer)
		*/
		newmsg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			connRunning = false
		} else {
			/*
				fmt.Println(buffer)
				fmt.Println(bytesRead)
			*/
			var msg Msg
			//color.Blue(string(buffer[0:bytesRead]))
			//msg = msg.createFromBytes([]byte(string(buffer[0:bytesRead])))
			msg = msg.createFromBytes([]byte(newmsg))
			messageHandler(connPeer, msg)
		}
	}
	//TODO add that if the peer closed is the p2p server, show a warning message at the peer
	log.Println("Peer: " + conn.RemoteAddr().String() + " connection closed")
	conn.Close()
	//TODO delete the peer from the outcomingPeersList
	deletePeerFromPeersList(connPeer, &outcomingPeersList)
	color.Yellow("peer deleted, current peerList:")
	printPeersList()
}
