package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	conn.Close()
}

func client() {
	// connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send message
	msg := "Formez vos battaillons!"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	conn.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
