package main

import (
	"fmt"
	"log"
	"net"
	"encoding/gob"
	handler "leatep/handler"
	// "time"
)
type Msg struct {
    M string
}


func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
	handler.ReceiveHandle()
	for {
    	msg := &Msg{}
    	dec.Decode(msg)
		
		if msg.M == "" {
			fmt.Println("Client offline")
			break
		}
    	fmt.Printf("Received: %v\n", msg.M);
	}
	conn.Close()
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":9000")
	Err(err)

	fmt.Println("Listening to 9000")

	for {
		conn, err := ln.Accept() // listen for clients
		Err(err)
		
		go handleConnection(conn)
	}
}
