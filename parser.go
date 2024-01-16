package main

import (
	"errors"
	"io"
	"log"
	"net"
)

type RespReader struct {
    currentPos int
    buffer     []byte
}

func NewRespReader(conn net.Conn) *RespReader {
	buf := make([]byte, 1024)
		
	_, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Println(err)
			return nil
		}
	}

    return &RespReader{
        currentPos: 0,
        buffer:     buf,
    }
}

func (reader *RespReader) ReadNext() (string, error) {
    for i := reader.currentPos; i < len(reader.buffer); i++ {
        if i+2 < len(reader.buffer) && string(reader.buffer[i:i+2]) == "\r\n" {
            cmd := string(reader.buffer[reader.currentPos:i])
            reader.currentPos = i+2
            return cmd, nil
        }
    }

    return "", errors.New("no more commands")
}