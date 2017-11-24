package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fatih/color"
)

type Msg struct {
	Type      string    `json:"type"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	PeersList PeersList `json:"peerslist"`
}

func messageHandler(conn net.Conn, msg Msg) {

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

		time.Sleep(1000 * time.Millisecond)
		updatePeersList(conn, msg.PeersList)
		propagatePeersList()
		printPeersList()
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
