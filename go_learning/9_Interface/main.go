package main

import (
	"fmt"
	"leatep/pdb"
)

type lol struct {
	name string
	num int
	num1 int
}

func main() {
	pdb.Psql_connect()
	result, err := pdb.QueryUnits("select * from Unit;")

	fmt.Println(result, err)

	q, w, e := 1, 2, 3
	changing(&q,&w,&e)

	// in := make([]interface{}, 10)
	// in[1] = "asd"
	// fmt.Println(in)
// 
	// w1 := lol{"asd"}
	// fmt.Println(w1)
}

func changing(w ...interface{}) {
	for _, t := range w {
		switch d := t.(type) {
		case *int:
			*d = 15
			fmt.Println(*d)
		}
	}
}