package main

import (
    "fmt"
    "log"
    "net"
    // "encoding/gob"
	"time"
	"leatep/handler"
)

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":9999") // start connection
	Err(err)
	fmt.Println("client started")

	go handler.SendHandler(conn)
	go handler.ReceiveHandler(conn)
	for i:=0; i<100; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}