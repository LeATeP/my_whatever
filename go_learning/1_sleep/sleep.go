package main

import (
	"fmt"
	"time"
)

// #include <unistd.h>
// #include <errno.h>
// int usleep(useconds_t usec);
// import "C"



func main () {
	start := time.Now()

	var x int = 0
	t := 1 * time.Millisecond

	for x < 1000000{
		x++
		time.Sleep(t)
	}
	
	duration := time.Since(start)
	fmt.Println(duration)
}
