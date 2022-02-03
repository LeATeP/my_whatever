package main

import (
	"fmt"
)

func main() {
	total := 0
	for i:=0; i< 1000; i++ {
		total = total + i * i
		fmt.Println("level:", i, "xp:", i * i, "total:", total)
	}
}