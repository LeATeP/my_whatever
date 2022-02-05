package handler

import (
	"fmt"
	"time"
	"net"
	"encoding/gob"
)

 
type msgFormat struct {
	M string
}

func ReceiveHandler(conn net.Conn) {
    dec := gob.NewDecoder(conn)
	for {
    	msg := &msgFormat{}
    	dec.Decode(msg)

		if msg.M == "" {
			fmt.Println("Client offline")
			break
		}
    	fmt.Printf("Received: %v\n", msg.M);
	}
	conn.Close()
}

func SendHandler(conn net.Conn) {
	encoder := gob.NewEncoder(conn) // open new thread for connection
	for n1 := 0; n1 < 10; n1++{
		
		msg := &msgFormat{"asd"}
    	encoder.Encode(msg) // send msg to server through encoder 
		
		time.Sleep(1 * time.Second)
	}
    conn.Close() // close connection / unnecessary 
}
