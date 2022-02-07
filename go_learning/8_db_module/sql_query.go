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

type Units struct {
	Id int
	Name, Fraction, State, Job string
	Level, Xp int
}
 
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

func QueryUnits(sql_cmd string) ([]Units, error) {
	var allUnits []Units

	rows, err := db.Query(sql_cmd)
	CheckError(err, "attempt to query db.Query")
	defer rows.Close()

	for rows.Next() {
		var unit Units
		err := rows.Scan(&unit.Id, &unit.Name, &unit.Fraction, &unit.State, &unit.Job, &unit.Level, &unit.Xp)
		CheckError(err, "attempt to Iter through rows.Next")

		allUnits = append(allUnits, unit)
		fmt.Println(unit.Id)
	}
	err = rows.Err()
	CheckError(err, "attempt ending rows.Err")

	return allUnits, nil
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
