package main

import (
	"fmt"
	"math/rand"
	"time"
)

type num int
type tw interface {
	random() num
}

func main() {
	rand.Seed(time.Now().Unix())

	var n num
	var a tw
	a = &n

	fmt.Println(a.random())
}

func (c *num) random() num {
	*c = num(rand.Int63n(1000))
	return *c
}