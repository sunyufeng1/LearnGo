package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":8086"

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	println(11)
	for {
		handleClient(conn)
	}

}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte

	len, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		return
	}
	println(string(buf[:len]))

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
