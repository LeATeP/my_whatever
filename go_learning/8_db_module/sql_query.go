package pdb

import (
    "database/sql"
    "fmt"
	"log"
	// "time"
    _ "github.com/lib/pq"
)

var db *sql.DB

const (
    host     = "localhost" // "postgres"
    port     = 5432
    user     = "postgres"
    password = "123"
    dbname   = "sql"
)

// os.Getenv("HOSTNAME") # as an example how to get env

func Psql_connect() {
        // connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

        // open database
	var err error
    db, err = sql.Open("postgres", psqlconn)
    CheckError(err, "failed open connection sql.Open")

        // check db
    err = db.Ping()
    CheckError(err, "failed ping")

    fmt.Println("Connected!")

}

func createPointers(rows *sql.Rows) ([][]string, []string, []string, []interface{}) {
	columns, err := rows.Columns()
	CheckError(err, "columns")

	sliceOfRows := [][]string{}
	content_slice := make([]string, len(columns))
	pointers := make([]interface{}, len(columns))
	for i := range content_slice {
		pointers[i] = &content_slice[i]
	}
	
	return sliceOfRows, content_slice, columns, pointers
}

func convetIntoMap(slices [][]string, columns []string) []map[string]string {
	newMaps := []map[string]string{}

	for _, data := range slices {
		newMap := map[string]string{}
		for i, colName := range columns {
			newMap[colName] = data[i]
		}
		newMaps = append(newMaps, newMap)
	}
	return newMaps
}


func QueryUnits(sql_cmd string) ([]map[string]string, error) {
	rows, err := db.Query(sql_cmd)
	CheckError(err, "attempt to query db.Query")
	defer rows.Close()

	slices, rowData, columns, pointers := createPointers(rows)

	for rows.Next() {
		err := rows.Scan(pointers...)
		CheckError(err, "attempt to Iter through rows.Next")
		
		slices = append(slices, rowData)
	}
	err = rows.Err()
	CheckError(err, "attempt ending rows.Err")

	formedMap := convetIntoMap(slices, columns)
	return formedMap, nil
}

func Exec(sql_cmd string) bool {
	_, err := db.Exec(sql_cmd)
	if err != nil {
		return false
	}
	return true
}

func Close() {
	db.Close()
}

func CheckError(err error, error_name string) {
	if err != nil {
		log.Fatal(err, "\n--- ", error_name, " ---")
	}
}
