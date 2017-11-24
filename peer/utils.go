package main

import (
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
