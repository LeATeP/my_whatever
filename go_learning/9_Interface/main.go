package main

import (
	"fmt"
)

type lol struct {
	name string
}

func main() {
	in := make([]interface{}, 10)
	in[1] = "asd"
	fmt.Println(in)

	w1 := lol{"asd"}
	fmt.Println(w1)
}