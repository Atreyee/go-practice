package main

import (
	"fmt"
	"net"
	"os"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error!!")
		os.Exit(-1)
	}
}

func main() {
	fmt.Println("Simple server")
	ln, err := net.Listen("tcp", ":7890")
	errorHandler(err)
	fmt.Println("Ready to accept connections")
	conn, err := ln.Accept()
	errorHandler(err)
	fmt.Println("We have a client")

	var buf [1024]byte
	n, err := conn.Read(buf[:])
	errorHandler(err)

	fmt.Println("client sent", string(buf[:n]))

	conn.Close()
}
