package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

type Msg struct {
	Type      string    `json:"type"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	PeersList PeersList `json:"peerslist"`
}

func messageHandler(peer Peer, msg Msg) {

	log.Println("[New msg]")
	log.Println(msg)

	switch msg.Type {
	case "Hi":
		color.Yellow(msg.Type)
		color.Green(msg.Content)
		break
	case "PeersList":
		color.Blue("newPeerslist")
		fmt.Println(msg.PeersList)

		//time.Sleep(1000 * time.Millisecond)
		updateNetworkPeersList(peer.Conn, msg.PeersList)
		propagatePeersList(peer)
		printPeersList()
		break
	case "PeersList_Response":
		color.Blue("newPeerslist")
		fmt.Println(msg.PeersList)

		//time.Sleep(1000 * time.Millisecond)
		updateNetworkPeersList(peer.Conn, msg.PeersList)
		printPeersList()
		break
	case "MyID":
		color.Blue("MyID")
		fmt.Println(msg.Content)
		color.Green(peer.Conn.RemoteAddr().String())
		peer.ID = msg.Content
		searchPeerAndUpdate(peer)

		//time.Sleep(1000 * time.Millisecond)
		/*
			updatePeersList(peer.Conn, msg.PeersList)
			propagatePeersList(peer)
		*/
		printPeersList()
		break
	default:
		log.Println("Msg.Type not supported")
		break
	}

}
func (msg Msg) construct(msgtype string, msgcontent string, peersList PeersList) Msg {
	msg.Type = msgtype
	msg.Content = msgcontent
	msg.PeersList = peersList
	msg.Date = time.Now()
	return msg
}
func (msg Msg) toBytes() []byte {
	msgS, err := json.Marshal(msg)
	check(err)
	l := string(msgS) + "\n"
	r := []byte(l)
	return r
}
func (msg Msg) createFromBytes(bytes []byte) Msg {
	err := json.Unmarshal(bytes, &msg)
	check(err)
	return msg
}
