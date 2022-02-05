package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func say_now(name int) {
	for i:=0; i<10; i++ {
	fmt.Printf("user: %v | count: %v\n", name, i)
	time.Sleep(1 * time.Second)
	
	}
}


func main() {
	name := []string{"good", "nice", "awsome"}
	for n := range name {
		go say_now(n)
	}
	time.Sleep(10 * time.Second)
}
