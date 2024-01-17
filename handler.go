/*
 * This file contains is responsible for carrying out parsed commands
 * and generating responses.
 */

package main

func handler(command Command) (string) {
	switch command.Type {
	case "ECHO":
		return buildString(command.Args[0])
	case "PING":
		return buildString("PONG")
	case "COMMAND":
		return buildString("PONG")
	default:
		return buildString("PONG")
	}
}

func buildString(msg string) string {
	return "+" + msg + "\r\n"
}

func buildArr(msg string) string {
	return "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
}
