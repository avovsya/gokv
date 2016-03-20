package main

import (
	"bufio"
	"github.com/avovsya/gokv/session"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	var status session.Status

	defer conn.Close()
	defer func() {
		// Closure to capture 'status' variable
		log.Printf("Connection closed: %s. Status: %v.", conn.RemoteAddr(), status)
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	log.Printf("Incoming connection from: %s", conn.RemoteAddr())

	session := session.NewSession(conn.RemoteAddr().String(), reader, writer)

	status = session.Talk()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	log.Println("Server listening on port 8080. Use telnet '127.0.0.1 8080' to connect")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
