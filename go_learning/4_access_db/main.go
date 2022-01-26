
package main
 
import (
    "database/sql"
    "fmt"
	"log"
	"time"
    _ "github.com/lib/pq"
)

var db *sql.DB

const (
    host     = "localhost"
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
 
func main() {
	// os.Getenv("HOSTNAME") # as an example how to get env

	establish_conn()
    defer db.Close() // close database


	updateItem()
	// items, err := getItems()
	// CheckError(err, "Failed exec getItems")
	// for _, value := range items {
		// fmt.Println(value.id, value.name, value.amount)
	// }
}

func establish_conn() {
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

func updateItem() {
	for {
		db.Exec("update items set amount = amount + 1 where id = 1;")
		fmt.Println("+1")
		time.Sleep(10 * time.Millisecond)
	}
}

func CheckError(err error, error_name string) {
    if err != nil {
        log.Fatal(err, "\n--- ", error_name, " ---")
    }
}
