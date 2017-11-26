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

var outcomingPeersList PeersList
var incomingPeersList PeersList
var networkPeersList PeersList //the peers that have been received in the lists from other peers

/*
a future option is to put:
	type PeersList struct {
		Incoming	PeersList
		Outcoming	PeersList
		Network	PeersList
	}
*/

func peerIsInPeersList(p Peer, pl []Peer) int {
	r := -1
	for i, peer := range pl {
		if peer.IP+":"+peer.Port == p.IP+":"+p.Port {
			r = i
		}
	}
	return r
}

func deletePeerFromPeersList(p Peer, pl *PeersList) {
	i := peerIsInPeersList(p, pl.Peers)
	if i != -1 {
		//delete peer from pl.Peers
		pl.Peers = append(pl.Peers[:i], pl.Peers[i+1:]...)
	}
}
func appendPeerIfNoExist(pl PeersList, p Peer) PeersList {
	i := peerIsInPeersList(p, pl.Peers)
	if i == -1 {
		pl.Peers = append(pl.Peers, p)
	}
	return pl
}
func updateNetworkPeersList(conn net.Conn, newPeersList PeersList) {
	for _, peer := range newPeersList.Peers {
		if getIPPortFromConn(conn) == peer.IP+":"+peer.Port {
			peer.ID = newPeersList.PeerID
			color.Yellow(peer.ID)
		}
		i := peerIsInPeersList(peer, networkPeersList.Peers)
		if i == -1 {
			networkPeersList.Peers = append(networkPeersList.Peers, peer)
		} else {
			fmt.Println(networkPeersList.Peers[i])
			networkPeersList.Peers[i].ID = peer.ID
		}
	}
}
func searchPeerAndUpdate(p Peer) {
	for _, peer := range outcomingPeersList.Peers {
		color.Red(p.IP + ":" + p.Port)
		color.Yellow(peer.IP + ":" + peer.Port)
		if p.IP+":"+p.Port == peer.IP+":"+peer.Port {
			peer.ID = p.ID
		}
	}
}

//send the outcomingPeersList to all the peers except the peer p that has send the outcomingPeersList
func propagatePeersList(p Peer) {
	for _, peer := range networkPeersList.Peers {
		if peer.Conn != nil {
			if peer.ID != p.ID && p.ID != "" {
				color.Yellow(peer.ID + " - " + p.ID)
				var msg Msg
				msg.construct("PeersList", "here my outcomingPeersList")
				msg.PeersList = outcomingPeersList
				msgB := msg.toBytes()
				_, err := peer.Conn.Write(msgB)
				check(err)
			} else {
				/*
					for the moment, this is not being called, due that in the IncomingPeersList,
					there is no peer.ID, so in the comparation wih the peer that has send the
					peersList, is comparing ID with "", so nevere enters this 'else' section

					maybe it's not needed. TODO check if it's needed the PeerList_Response
					For the moment is working without it
				*/
				//to the peer that has sent the peerList, we send our PeersList

				var msg Msg
				msg.construct("PeersList_Response", "here my outcomingPeersList")
				msg.PeersList = outcomingPeersList
				msgB := msg.toBytes()
				_, err := peer.Conn.Write(msgB)
				check(err)
			}
		} else {
			//connect to peer
			if peer.ID != p.ID && peer.ID != runningPeer.ID {
				if peerIsInPeersList(peer, outcomingPeersList.Peers) == -1 {
					color.Red("no connection, connecting to peer: " + peer.Port)
					connectToPeer(peer)
				}
			}

		}
	}
}
func printPeersList() {
	fmt.Println("")
	color.Blue("runningPeer.ID: " + runningPeer.ID)
	color.Green("OUTCOMING PEERSLIST:")
	for _, peer := range outcomingPeersList.Peers {
		fmt.Println(peer)
	}
	color.Green("INCOMING PEERSLIST:")
	for _, peer := range incomingPeersList.Peers {
		fmt.Println(peer)
	}

	color.Green("NETWORK PEERSLIST:")
	for _, peer := range networkPeersList.Peers {
		fmt.Println(peer)
	}
	fmt.Println("")
}

//send the block to all the peers of the outcomingPeersList
func propagateBlock(b Block) {
	//prepare the msg to send to all connected peers
	var msg Msg
	msg.construct("Block", "new block")
	msg.Block = b
	msgB := msg.toBytes()
	for _, peer := range outcomingPeersList.Peers {
		if peer.Conn != nil {
			_, err := peer.Conn.Write(msgB)
			check(err)
		}
	}
}
