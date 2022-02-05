package main

import (
	"fmt"
)

func main() {
	// xp req for base levels 
	xp := 10000
	for i:=2; i<11; i++ {
		fmt.Printf("level %v: %v\n", i, xp)
		xp = xp * 2
	}



	// xp req for level, algorithm
	// total := 0
	// for i:=0; i< 1000; i++ {
		// total = total + i * i
		// fmt.Println("level:", i, "xp:", i * i, "total:", total)
	// }
}