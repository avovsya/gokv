package main

import (
	"bufio"
	"github.com/avovsya/gokv/session"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	var err error
	defer conn.Close()
	defer func() {
		// Closure to capture 'err' variable
		log.Printf("Connection closed: %s. Error: %v", conn.RemoteAddr(), err)
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	log.Printf("Incoming connection from: %s", conn.RemoteAddr())

	session := session.NewSession(conn.RemoteAddr().String(), reader, writer)

	session.Ready()
	err = session.Talk()
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
