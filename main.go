package main

import (
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	
	for {
		reader := NewRespReader(conn)
		
		cmd, err := reader.ReadNext()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(cmd)

		if cmd[0] != '*' {
			conn.Write([]byte("-ERR unknown command '" + cmd + "'\r\n"))
			return
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