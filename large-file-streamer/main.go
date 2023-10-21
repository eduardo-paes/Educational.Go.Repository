package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)
const PORT string = ":8080"
const TRANSFER_RATE int64 = 1024
const TRANSFER_TYPE string = "tcp"

type FileServer struct{}

// Start the server
func (fs * FileServer) start() {
	ln, err := net.Listen(TRANSFER_TYPE, PORT)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go fs.readLoop(conn)
	}
}

// Read received file loop
func (fs * FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		// Read the size of the file
		var size int64
		var sizeToRead int64 = TRANSFER_RATE
		binary.Read(conn, binary.LittleEndian, &size)

		for size > 0 {
			// If the remaining size is less than the bit rate, read the remaining size
			if size < TRANSFER_RATE {
				sizeToRead = size
			}

			// Read the file
			n, err := io.CopyN(buf, conn, sizeToRead)
			if err != nil {
				log.Fatal(err)
			}

			// Print the file
			fmt.Println(buf.Bytes())
			fmt.Printf("Read %d bytes\n", n)

			// Reset the buffer
			size -= int64(sizeToRead)
			buf.Reset()
		}
	}	
}

// Send a file to the server
func sendFile(size int) error {
	file := make([]byte, size)

	// Generate a random file
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	// Connect to the server
	conn, err := net.Dial(TRANSFER_TYPE, PORT)
	if err != nil {
		return err
	}

	// Write the size of the file
	binary.Write(conn, binary.LittleEndian, int64(size))

	// Write the file
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}

	// Print the number of bytes written
	fmt.Printf("Written %d bytes\n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		err := sendFile(123456)
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	server := &FileServer{}
	server.start()
}