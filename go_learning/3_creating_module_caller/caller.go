package main

import (
	"fmt"
	"leatep/creating_module"
	"log"
)

func main() {
	log.SetPrefix("greeting: ") // 1. at the start of the error msg 
	log.SetFlags(0) // 2. set log, settings

	msg, err := creating_module.Hello("asd")
	if err != nil { // catching error
		log.Fatal(err)
	}
	fmt.Println(msg)	

	names := []string {
		"1asd",
		"2qwe",
		"3zxc",
	}
	vars, err := creating_module.FormatList(names)
	if err != nil {log.Fatal(err)}

	for i, v := range vars {
		fmt.Println(i, v) }

}

