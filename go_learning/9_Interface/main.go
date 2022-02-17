package main

import (
	"fmt"
	"leatep/pdb"
	"math/rand"
	"time"
)

func main() {
	pdb.Psql_connect()
	result, _ := pdb.QueryUnits("select id, name, amount from items;")

	for _, row := range result {
		fmt.Printf("%v: %v, %v\n", row["id"], row["name"], row["amount"])
	}
	rand.Seed(time.Now().UnixNano())


	t_old, w, count := 0, 0, 0
	for {
		t_new := time.Now().Second()
		if t_old != t_new {
			t_old = t_new
			fmt.Println(t_new, count, w)
			count = 0
		}

		w = rand1()
		count = count +1
		pdb.Exec("update items set amount = amount +1 where id = 2")
	}
}

func rand1() int {
	r, w := 0, 0
	for i:=0; i < 1000000; i++ {
		r = rand.Intn(10000000)
		w = w+r
	}
	return w
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
