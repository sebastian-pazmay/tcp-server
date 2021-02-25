package main

import (
	"flag"
	"fmt"
	"net"
)

var netAddress string
var netPort string
var length = 0

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.StringVar(&netAddress, "a", "localhost", "Specify IP address. Default is localhost")
	flag.StringVar(&netPort, "p", "9090", "Specify network port. Default is 9090")
	flag.Parse()
	netSocket := netAddress + ":" + netPort
	tcpListener, err := net.Listen("tcp", netSocket)
	CheckError(err)
	fmt.Println("I am listening to TCP on: ", netSocket)
	for {
		tcpConnection, err := tcpListener.Accept()
		CheckError(err)
		fmt.Println("TCP connection accepted from: ", tcpConnection.RemoteAddr())
		go handleTCPRequest(tcpConnection)
	}
}

func handleTCPRequest(connection net.Conn) {
	rcvBuffer := make([]byte, 104857600)
	data, err := connection.Read(rcvBuffer)
	CheckError(err)
	if data != 0 {
		length = data + length
		fmt.Printf("READ %d bytes\n", length)
		_, err := connection.Write([]byte("TCP Server response"))
		CheckError(err)
		fmt.Println("Server responded!")
	}
	err = connection.Close()
	CheckError(err)
}
