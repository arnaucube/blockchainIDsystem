package main

import (
	"encoding/json"
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
		listPeers = append(listPeers, newPeer)
		go handleConn(conn)
	}
}
func connectToPeer(peer Peer) {
	conn, err := net.Dial("tcp", peer.IP+":"+peer.Port)
	if err != nil {
		log.Println("Error connecting to: " + peer.IP + ":" + peer.Port)
		return
	}
	peer.Conn = conn
	listPeers = append(listPeers, peer)
	go handleConn(conn)
}
func handleConn(conn net.Conn) {
	connRunning := true
	log.Println("handling conn: " + conn.RemoteAddr().String())
	//reply to the conn
	msgB := newMsgBytes("Hi", "New Peer connected")
	_, err := conn.Write(msgB)
	if err != nil {
		check(err)
	}

	buffer := make([]byte, 1024)
	for connRunning {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			connRunning = false
		} else {
			s := string(buffer[0:bytesRead])
			var msg Msg
			err := json.Unmarshal([]byte(s), &msg)
			check(err)
			log.Println("[New msg] [Title]: " + msg.Type + " [Content]: " + msg.Content)
			color.Green(msg.Content)
		}
	}
	//TODO add that if the peer closed is the p2p server, show a warning message at the peer
	log.Println("Peer: " + conn.RemoteAddr().String() + " connection closed")
	conn.Close()
}
