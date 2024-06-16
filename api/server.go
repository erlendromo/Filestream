package api

import (
	"net"
)

type Server interface {
	Serve()
	readLoop(conn net.Conn)
}

func guard(err error) {
	if err != nil {
		panic(err)
	}
}
