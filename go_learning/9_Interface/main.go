package main

import (
	"fmt"
	"leatep/pdb"
)

type lol struct {
	num int
	num1 int
}

func main() {
	pdb.Psql_connect()
	result, err := pdb.QueryUnits("select * from unit;")
	fmt.Println(result, err)

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
