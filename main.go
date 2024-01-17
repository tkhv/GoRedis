package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Listening on port 6379.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
        if r := recover(); r != nil {
            log.Println(conn.RemoteAddr(), " terminated.")
        }
    }()

	defer conn.Close()
	
	log.Println("New connection from ", conn.RemoteAddr())
	for {
		reader := parseResp(conn)
		log.Println(reader)
		conn.Write([]byte("+PONG\r\n"))
	}
}