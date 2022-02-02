package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
	"time"
)

type P struct {
    M string
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":9000")
	Err(err)
	
	fmt.Println("client started")

	encoder := gob.NewEncoder(conn)
	for n1 := 0; n1 < 10; n1++{
		
		p := &P{fmt.Sprint(n1)}
    	encoder.Encode(p)
		
		time.Sleep(1 * time.Second)
	}
    conn.Close()
    fmt.Println("done")
}