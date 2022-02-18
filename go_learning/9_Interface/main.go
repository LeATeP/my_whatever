package main

import (
	"fmt"
	"leatep/pdb"
)

func main() {
	pdb.Psql_connect()
	result, _ := pdb.QueryUnits("select id, name, amount from items;")

	for _, row := range result {
		fmt.Printf("%v: %v, %v\n", row["id"], row["name"], row["amount"])
	}

}

