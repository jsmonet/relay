package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	inPort  = flag.Int("in", 1, "program will listen on this port and send traffic to the port you declare for out")
	outPort = flag.Int("out", 1, "This is the port an active service is running on. The port you specify for 'in' will ship traffic to a process running on this port")
)

func passAlong(conn net.Conn, destPort string) {
	client, err := net.Dial("tcp", destPort)
	if err != nil {
		log.Fatalf("net.Dial failed: %v", err)
	}
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(conn, client)
	}()
}

func main() {
	flag.Parse()

	if *inPort == 1 || *outPort == 1 {
		fmt.Println("You did not specify either/both an in port and an out port")
		os.Exit(1)
	}

	var soakAddress string
	var passAddress string

	soakAddress = fmt.Sprintf("0.0.0.0:%v", *inPort)
	passAddress = fmt.Sprintf("127.0.0.1:%v", *outPort)

	soak, err := net.Listen("tcp", soakAddress)
	if err != nil {
		log.Fatalf("Listener failed to init: %v", err)
	}

	for {
		conn, err := soak.Accept()
		if err != nil {
			log.Fatalf("Error: couldn't accept the listener. Here's the error: %v", err)
		}
		go passAlong(conn, passAddress)
	}
}
