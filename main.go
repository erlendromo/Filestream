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
	file, err := os.Open("./jaja.pdf")
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	size := stat.Size()
	binary.Write(conn, binary.LittleEndian, size)

	n, err := io.CopyN(conn, file, size)
	if err != nil {
		return err
	}

	file.Close()

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
