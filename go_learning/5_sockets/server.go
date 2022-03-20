package main

import (
	"fmt"
	"log"
	"net"
	"encoding/json"
	// "encoding/gob"
	"leatep/handler"
	"leatep/pdb"
	// "time"
)

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sliceToMap(item interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(item)
	json.Unmarshal(j, &m)
	
	return &m
}


func main() {
	ln, err := net.Listen("tcp", ":9000")
	Err(err)

	fmt.Println("Listening to 9000")
	pdb.Psql_connect()
	result, err := pdb.QueryUnits("select * from unit;")
	if err != nil {
		fmt.Println("failed query")
	}
	for _, r := range result { 
		item := sliceToMap(r)
		fmt.Println(item)
		
	}


	for {
		conn, err := ln.Accept() // listen for clients
		Err(err)
		
		go handler.ReceiveHandler(conn)
		go handler.SendHandler(conn)
	}
}
