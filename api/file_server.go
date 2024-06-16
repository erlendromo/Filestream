package api

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type FileServer struct {
	Port string
}

func NewFileServer(p string) Server {
	return &FileServer{
		Port: p,
	}
}

func (fs *FileServer) Serve() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", fs.Port))
	guard(err)

	for {
		conn, err := ln.Accept()
		guard(err)

		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	buffer := new(bytes.Buffer)
	var size int64

	binary.Read(conn, binary.LittleEndian, &size)

	n, err := io.CopyN(buffer, conn, size)
	guard(err)

	fmt.Println(buffer.Bytes())
	fmt.Printf("Recieved %d bytes over the network\n", n)
}
