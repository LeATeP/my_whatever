package pdb

import (
    "database/sql"
    "fmt"
	"log"
    _ "github.com/lib/pq"
)

var db *sql.DB

const (
    host     = "postgres" // "postgres"
    port     = 5432
    user     = "postgres"
    password = "123"
    dbname   = "sql"
)

func Psql_connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							host, port, user, password, dbname)
	var err error
    db, err = sql.Open("postgres", psqlconn) // connecting to db
    CheckError(err, "failed open connection sql.Open")

    err = db.Ping() // check connection
    CheckError(err, "failed ping")
    fmt.Println("Connected!")
}

func createPointers(content *[]string) ([]interface{}) {
	point 		:= make([]interface{}, len(*content))
	for i		:= range *content {
		point[i] = &(*content)[i]
	}
	return point
}

func convetIntoMap(slices [][]string, columns []string) []map[string]string {
	newMaps 	:= []map[string]string{}

	for _, data := range slices {
		newMap 	:= map[string]string{}
		for i, colName := range columns {
			newMap[colName] = data[i]
		}
		newMaps = append(newMaps, newMap)
	}
	return newMaps
}


func QuerySelect(sql_cmd string) ([]map[string]string, error) {
	rows, err 		:= db.Query(sql_cmd)
	CheckError(err, "attempt to query db.Query")
	defer rows.Close()

	columns, _ 		:= rows.Columns()
	rowsStack		:= [][]string{}
	for rows.Next() {
		content 	:= make([]string, len(columns))
		pointers 	:= createPointers(&content)
	
		err := rows.Scan(pointers...)
		CheckError(err, "attempt to Scan rows.Next")
		
		rowsStack 	= append(rowsStack, content)
	}
	err = rows.Err()
	CheckError(err, "attempt ending rows.Err")

	formedMap 		:= convetIntoMap(rowsStack, columns)
	return formedMap, nil
}

func Exec(sql_cmd string) bool {
	_, err := db.Exec(sql_cmd)
	return err == nil
}

func Close() {
	db.Close()
}

func CheckError(err error, error_name string) {
	if err != nil {
		log.Fatal(err, "\n--- ", error_name, " ---")
	}
}
