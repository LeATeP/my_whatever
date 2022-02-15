package main

import (
	"fmt"
	"leatep/pdb"
)

func main() {
	pdb.Psql_connect()
	result, _ := pdb.QueryUnits("select * from items;")

	for _, t := range result {
		fmt.Println(t)
	}
}

func changing(w ...interface{}) {
	for _, t := range w {
		switch d := t.(type) {
		case *int:
			fmt.Println(*d)
			*d = 15
		case *string:
			*d = "asd"
		}
	}
}
