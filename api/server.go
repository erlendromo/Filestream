package api

import (
	"log"
	"net"
)

// Interface of the server to allow mocking and/or testing
type Server interface {
	Serve()
	readLoop(conn net.Conn)
}

// Guard-clause to ensure that critical errors stops, and recovery can happen
func guard(err error) {
	if err != nil {
		panic(err)
	}
}

func recovery() {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic: %v", r)
	}
}
