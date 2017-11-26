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
	Block     Block     `json:"block"`
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
		color.Red("PeersList")

		updateNetworkPeersList(peer.Conn, msg.PeersList)
		propagatePeersList(peer)
		printPeersList()
		break
	case "PeersList_Response":
		//for the moment is not beeing used
		color.Blue("newPeerslist, from PeersList_Response")
		fmt.Println(msg.PeersList)
		color.Red("PeersList_Response")

		updateNetworkPeersList(peer.Conn, msg.PeersList)
		propagatePeersList(peer)
		printPeersList()
		break
	case "Block":
		//TODO check if the block is signed by an autorized emitter
		if !blockchain.blockExists(msg.Block) {
			blockchain.addBlock(msg.Block)
			propagateBlock(msg.Block)
		}
		break
	default:
		log.Println("Msg.Type not supported")
		break
	}

}
func (msg *Msg) construct(msgtype string, msgcontent string) {
	msg.Type = msgtype
	msg.Content = msgcontent
	msg.Date = time.Now()
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
