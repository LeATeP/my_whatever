package main
import (
	"fmt"
	"log"
	"net"
	"encoding/gob"
	// "time"
)
type Msg struct {
    M string
}


func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
	for {
    	msg := &Msg{}
    	dec.Decode(msg)
		
		if msg.M == "" {
			fmt.Println("Client offline")
			conn.Close()
			break
		}
    	fmt.Printf("Received: %v\n", msg.M);
	}
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
