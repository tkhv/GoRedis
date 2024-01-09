package main

import (
	"fmt"
	"io"
	"net"
)

func checkErr(err error){
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":6379")
	checkErr(err)
	fmt.Println("Listening on port 6379.")

	conn, err := ln.Accept()
	checkErr(err)
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
				panic(err)
			}
		}
		conn.Write([]byte("+JENKINS DEPLOYED THIS FROM EC2\r\n"))
	}
}