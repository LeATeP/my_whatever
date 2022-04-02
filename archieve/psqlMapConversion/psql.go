// package psql is mainly to give interface to manage psql db query's and stuff
package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// dbStruct is to be used as a `connection` handler, for method to use
type dbStruct struct {
	conn *sql.DB
}

// DbInterface to package all methods and to be accessible and understanable
type DbInterface interface {
	Exec(string) error
	Ping() error
	QuerySelect(string) ([]map[string]any, error)
}

var (
	d dbStruct
	err error
)

// Psql_connect is main initialize fn, that connect to db and give interface
func Psql_connect() (DbInterface, error) {
	config 		:= init_config()
	psqlconn 	:= fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							config.host, config.port, config.user, config.password, config.dbname)
	
	// Connecting to db
    d.conn, err = sql.Open("postgres", psqlconn)			
	if err	   	!= nil { return nil, err }

	// checking connection if working
    err 		= d.Ping(); 									
	if err 		!= nil { return nil, err }

	return DbInterface(&d), nil
}

// QuerySelect is request data from db, and reformat it into accessible Map format
func (d *dbStruct) QuerySelect(sql_cmd string) ([]map[string]any, error) {
	rows, err 		:= d.conn.Query(sql_cmd)
	if err != nil {	return nil, err }
	defer rows.Close()

	columns, err 	:= rows.Columns()
	if err != nil {	return nil, err }

	rowsStack, err  := iterRows(rows, len(columns))
	if err != nil {	return nil, err }

	formedMap 		:= convetIntoMap(rowsStack, columns)
	return formedMap, nil
}
// iterRows help to reformat data, that came from request in QuerySelect
func iterRows (rows *sql.Rows, row_length int) ([][]any, error) {
	rowsStack		:= [][]any{}

	for rows.Next() {
		content, pointers := makePointers(row_length)

		err := rows.Scan(pointers...)
		if err != nil {	return nil, err }
		
		rowsStack 	= append(rowsStack, content)
	}
	err = rows.Err()
	if err != nil {	return nil, err }
	
	return rowsStack, nil
}
// Exec is a main func to exec commands in sql that isn't required output
// like update, insert, delete
func (d *dbStruct) Exec(sql_cmd string) error {
	_, err := d.conn.Exec(sql_cmd)					
	if err != nil { return err }
	return nil
}
// Ping to check if connection to db is stable
func (d *dbStruct) Ping() error {
	return d.conn.Ping()
}