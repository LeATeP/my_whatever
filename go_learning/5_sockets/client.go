package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
	"time"
	handler "leatep/handlers"
)

type P struct {
    M string
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnection(conn net.Conn) {
	encoder := gob.NewEncoder(conn) // open new thread for connection
	for n1 := 0; n1 < 10; n1++{
		
		p := &P{fmt.Sprint(n1)}
    	encoder.Encode(p) // send msg to server through encoder 
		
		time.Sleep(1 * time.Second)
	}
    conn.Close() // close connection / unnecessary 
}

func main() {
	conn, err := net.Dial("tcp", ":9000") // start connection
	Err(err)
	fmt.Println("client started")

	go handleConnection(conn)
	for i:=0; i<100; i++ {
		fmt.Println(i)
	}
}