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

	for x < 1000000{
		x++
		// C.usleep(1000)
		time.Sleep(1 * time.Millisecond)
	}
	
	duration := time.Since(start)
	fmt.Println(duration)
}
