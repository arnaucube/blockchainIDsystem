package main

import (
	"encoding/json"
	"net"
	"strings"
)

func newMsgBytes(msgtype string, msgcontent string) []byte {
	var msg Msg
	msg.Type = msgtype
	msg.Content = msgcontent
	msgS, err := json.Marshal(msg)
	check(err)
	return msgS
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
