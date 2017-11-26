package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net"
	"strings"
)

func getIPPortFromConn(conn net.Conn) string {
	ip := getIPFromConn(conn)
	port := getPortFromConn(conn)
	return ip + ":" + port
}
func getIPFromConn(conn net.Conn) string {
	s := conn.RemoteAddr().String()
	s = strings.Split(s, ":")[0]
	s = strings.Trim(s, ":")
	return s
}
func getPortFromConn(conn net.Conn) string {
	s := conn.RemoteAddr().String()
	s = strings.Split(s, ":")[1]
	s = strings.Trim(s, ":")
	return s
}
func randInt(min int, max int) int {
	r := rand.Intn(max-min) + min
	return r
}
func hashPeer(p Peer) string {
	peerString := p.IP + ":" + p.Port

	h := sha256.New()
	h.Write([]byte(peerString))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
func hashBlock(b Block) string {
	blockJson, err := json.Marshal(b)
	check(err)
	blockString := string(blockJson)

	h := sha256.New()
	h.Write([]byte(blockString))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
