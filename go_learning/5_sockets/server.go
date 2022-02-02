package main
import (
	"fmt"
	"log"
	"net"
	"encoding/gob"
	"time"
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
			break
		}
		time.Sleep(500 * time.Millisecond)
		// if  {
			// log.Fatal("Connection is broken")
		// }
    	fmt.Printf("Received: %v\n", msg.M);
	}
	// conn.Close()
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
