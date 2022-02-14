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

	e := make([]interface{}, 3)
	r := make([]interface{}, len(e))
	for i := range e {
		r[i] = &e[i]
	}

	// w := make([]interface{}, 3)
	changing(r...)
	fmt.Println(e)

}

func changing(w ...interface{}) {
	for _, t := range w {
		switch d := t.(type) {
		case *int:
			fmt.Println(*d)
			*d = 15
		default:
			fmt.Println("nil")
		}
	}
}
