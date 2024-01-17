package main

import (
	"io"
	"log"
	"net"
	"strconv"
)
type Commands struct {
	count int
	cmds []Command
}

type Command struct {
	Type string
	Args []string
}

type RespReader struct {
    currentPos int
    buffer     []byte
	arrLen		int
}

/*
 * Parses a RESP request and returns a Commands struct containing
 * the number of commands and the commands themselves.
*/
func parseResp(conn net.Conn) (Commands) {
	reader := NewRespReader(conn)
	commands := Commands{count: 0, cmds: []Command{}}

	for {
		element, err := reader.NextElement()
		if err == io.EOF {
			break
		}

		if element == "ECHO" {
			arg, err := reader.NextElement()
			if err != nil {
				log.Panic("Not enough arguments for ECHO")
			}
			commands.cmds = append(commands.cmds, Command{Type: element, Args: []string{arg}})
			commands.count++
		} else {
			// Catch all
			commands.cmds = append(commands.cmds, Command{Type: element})
			commands.count++
		}
	}

	return commands
}

/* 
 * Returns a RespReader after looking for an array length.
 * Can panic if the request is malformed.
*/
func NewRespReader(conn net.Conn) *RespReader {
	buf := make([]byte, 1024)
		
	_, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Panic(err)
		}
	} else if buf[0] != '*' {
		log.Panic("Expected *, got ", buf[0])
	}

	i := 1
    for ; buf[i] != '\r'; i++ {}
    arrLen, _ := strconv.Atoi(string(buf[1:i]))

    return &RespReader{
        currentPos: i+2, // Skip CRLF
        buffer:     buf,
		arrLen:		arrLen,
    }
}

/*
 * Returns the next element in the request.
 * Will return io.EOF if a CRLF is not reached.
*/
func (reader *RespReader) NextElement() (string, error) {
	i, err := reader.FindCRLF()
	if err != nil {
		return "", err
	}
	reader.currentPos = i+2
	i, err = reader.FindCRLF()
	if err != nil {
		return "", err
	}
	str := string(reader.buffer[reader.currentPos:i])
	reader.currentPos = i+2
	return str, nil
}

/*
 * Returns the index of the next CRLF.
 * Will return io.EOF if a CRLF is not reached.
*/
func (reader *RespReader) FindCRLF() (int, error) {
	// Advance i until CRLF, return buf[currentPos:i] and update currentPos to i+2
	for i := reader.currentPos; i < len(reader.buffer); i++ {
		if reader.buffer[i] == '\r'  {
			return i, nil
		}
	}

	return 0, io.EOF
}