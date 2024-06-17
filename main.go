package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/erlendromo/Filestream/api"
)

func sendFile() error {
	// Establish tcp-connection to appropriate port
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		return err
	}

	defer conn.Close()

	// Read file
	file, err := os.Open("./jaja.pdf")
	if err != nil {
		return err
	}

	defer file.Close()

	// Get stats from file
	stat, err := file.Stat()
	if err != nil {
		return err
	}

	// Read the size of the file, and send it over the connection
	size := stat.Size()
	binary.Write(conn, binary.LittleEndian, size)

	// io.CopyN to stream the file over the connection, allocating proper size to the buffer
	n, err := io.CopyN(conn, file, size)
	if err != nil {
		return err
	}

	fmt.Printf("Sent %d bytes over the network\n", n)
	return nil
}

func main() {

	go func() {
		time.Sleep(3 * time.Second)
		sendFile()
	}()

	server := api.NewFileServer("8080")
	server.Serve()
}
