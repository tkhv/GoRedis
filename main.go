package main

import (
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	
	for {
		buf := make([]byte, 1024)
		
		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return
			}
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println(err)
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