package main

import (
	"fmt"
	"log"
	"net"
	// "encoding/gob"
	"leatep/handler"
	// "time"
)
type msg struct {
    M string
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
		
		go handler.ReceiveHandler(conn)
		go handler.SendHandler(conn)
	}
}
