package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.String("port", "80", "The specified port")
	flag.Parse()
	fmt.Printf("port: %s", *port)
	listener, err := net.Listen("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn1(conn)	// handle one connection at a time
	}
}

func handleConn1(c net.Conn)  {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return	// client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}