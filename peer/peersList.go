package main

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

type PeersList struct {
	PeerID string
	Peers  []Peer    `json:"peerslist"`
	Date   time.Time `json:"date"`
}

//var peersList []Peer
var peersList PeersList

func peerIsInPeersList(p Peer, pl []Peer) int {
	r := -1
	for i, peer := range pl {
		if peer.IP+":"+peer.Port == p.IP+":"+p.Port {
			r = i
		}
	}
	return r
}
func updatePeersList(conn net.Conn, newPeersList PeersList) {
	for _, peer := range newPeersList.Peers {
		if getIPPortFromConn(conn) == peer.IP+":"+peer.Port {
			peer.ID = newPeersList.PeerID
		}
		i := peerIsInPeersList(peer, peersList.Peers)
		if i == -1 {
			peersList.Peers = append(peersList.Peers, peer)
		} else {
			fmt.Println(peersList.Peers[i])
			peersList.Peers[i].ID = peer.ID
		}
	}
}

//send the peersList to all the peers except the peer that has send the peersList
func propagatePeersList() {
	for _, peer := range peersList.Peers {
		if peer.Conn != nil {
			var msg Msg
			msg = msg.construct("PeersList", "here my peersList", peersList)
			msgB := msg.toBytes()
			_, err := peer.Conn.Write(msgB)
			check(err)
		}
	}
}
func printPeersList() {
	fmt.Println("")
	color.Green("PEERSLIST:")
	color.Green("thisPeerId: " + thisPeerID)
	for _, peer := range peersList.Peers {
		fmt.Println(peer)
	}
	fmt.Println("")
}
