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
			color.Yellow(peer.ID)
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
func searchPeerAndUpdate(p Peer) {
	for _, peer := range peersList.Peers {
		color.Red(p.IP + ":" + p.Port)
		color.Yellow(peer.IP + ":" + peer.Port)
		if p.IP+":"+p.Port == peer.IP+":"+peer.Port {
			peer.ID = p.ID
		}
	}
}

//send the peersList to all the peers except the peer that has send the peersList
func propagatePeersList(p Peer) {
	for _, peer := range peersList.Peers {
		if peer.Conn != nil {
			if peer.ID != p.ID && p.ID != "" {
				color.Yellow(peer.ID + " - " + p.ID)
				var msg Msg
				msg = msg.construct("PeersList", "here my peersList", peersList)
				msgB := msg.toBytes()
				_, err := peer.Conn.Write(msgB)
				check(err)
			} else {
				//to the peer that has sent the peerList, we send our ID
				/*
					var msg Msg
					var pl PeersList
					msg = msg.construct("MyID", runningPeer.ID, pl)
					msgB := msg.toBytes()
					_, err := p.Conn.Write(msgB)
					check(err)
				*/
				var msg Msg
				msg = msg.construct("PeersList_Response", "here my peersList", peersList)
				msgB := msg.toBytes()
				_, err := peer.Conn.Write(msgB)
				check(err)
			}
		}
	}
}
func printPeersList() {
	fmt.Println("")
	color.Green("PEERSLIST:")
	color.Green("runningPeer.ID: " + runningPeer.ID)
	for _, peer := range peersList.Peers {
		fmt.Println(peer)
	}
	fmt.Println("")
}
