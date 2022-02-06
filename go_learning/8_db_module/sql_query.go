package dbHandler
 
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

type Item struct {
	id string
	name string
	amount int
}
 
func Exec(cmd string) (result bool, err error) {
	// os.Getenv("HOSTNAME") # as an example how to get env

	db1 := psql_connect()
    defer db1.Close() // close database

	total := 0
	for i:=2; i<1000; i++ {
		total += i*i
		updateItem(i, i*i, total, "mortal")
	}
	return
}

func psql_connect() (db *sql.DB) {
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
	return db

}

func getItems() ([]Item, error) {
	var items []Item

	rows, err := db.Query("Select id, name, amount from items order by id;")
	CheckError(err, "attempt to query db.Query")
	defer rows.Close()

	for rows.Next() {
		var itm Item
		err := rows.Scan(&itm.id, &itm.name, &itm.amount)
		CheckError(err, "attempt to Iter through rows.Next")

		items = append(items, itm)
	}
	err = rows.Err()
	CheckError(err, "attempt ending rows.Err")

	return items, nil
}

func updateItem(level, xp, total_xp int, grade string) {
	sql_cmd := fmt.Sprintf("insert into levels (level, grade, req_xp, total_xp) values (%v, '%v', %v, %v);", level, grade, xp, total_xp)
	
	_, err := db.Exec(sql_cmd)
	CheckError(err, "attemping to updateItem...:",)
	
}

func CheckError(err error, error_name string) {
    if err != nil {
        log.Fatal(err, "\n--- ", error_name, " ---")
    }
}
