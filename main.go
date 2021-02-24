package main

import (
	"fmt"
	"net"
)

// TODO:
// include read addr as argument
// use localhost by default
//const addr = "10.0.0.6:9090"
const addr = "localhost:9090"

var length = 0

// TODO:
// times is not required now, remove it
var times = 0

func CheckError(err error) {
	if err != nil {
		// TODO:
		// Fix error catching
		//fmt.Println("Error: " , err)
	}
}

func main() {
	// TODO:
	// Fix messages
	fmt.Println("This is the ultra new TCP server!")
	listener, err := net.Listen("tcp", addr)
	fmt.Println("I am listening to TCP on ", addr)
	CheckError(err)
	defer listener.Close()
	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		CheckError(err)
		fmt.Println("I accepted a TCP connection from ", conn.RemoteAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	for {
		buf := make([]byte, 10485760)
		// TODO:
		// Search what it is being done here
		n, err := conn.Read(buf)
		if n != 0 {
			//fmt.Println("\n Received ", string(buf[0:n]), " from ", conn.RemoteAddr())
			CheckError(err)
			// TODO:
			// Remove times
			times = times + 1
			length = n + length
			// log the bytes written
			fmt.Printf("READ %d bytes\n", length)
			// TODO:
			// Remove times
			fmt.Printf("TIMES %d times\n", times)
			conn.Write([]byte("TCP Server response!"))
			fmt.Println("Server also responded!")
		}
	}
}
