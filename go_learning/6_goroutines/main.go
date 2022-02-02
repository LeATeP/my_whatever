package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func say_now() {
	fmt.Println("starting")
	fmt.Printf("asd")
	time.Sleep(1 * time.Second)
}


func main() {
	name := []string{"good", "nice", "awsome"}
	for n := range name {
		fmt.Println(name[n])
		go say_now()
	}
}
