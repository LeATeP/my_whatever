// package psql utils.go is for fn used in psql
// but not necessary to be accessible outside of the package
package psql

import (
	"os"
	"sort"
)

// con_config is connection config
type con_config struct {
	hostname string
    host     string
    dbname   string
    user     string
    password string
	port     int
}

func init_config() *con_config {
	return &con_config{
		hostname:   os.Getenv("HOSTHAME"),
		host: 		os.Getenv("PSQL_HOST"),
  		dbname: 	os.Getenv("PSQL_DB"),
  		user: 		os.Getenv("PSQL_USER"),
  		password:   os.Getenv("PGPASSWORD"),
		port: 		5432,
	}
}

func convetIntoMap(slices [][]any, columns []string) []map[string]any {
	newMaps 	:= make([]map[string]any, len(slices))

	for i, data := range slices {
		newMap 	:= map[string]any{}
		for r, colName := range columns {
			newMap[colName] = data[r]
		}
		newMaps[i] = newMap
	}
	sortSliceOfMap(newMaps)
	return newMaps
}

func sortSliceOfMap(newMaps []map[string]any) {
	_, exist := newMaps[0]["id"].(int64)
	if exist {
  		sort.Slice(newMaps, func(i, j int) bool { return newMaps[i]["id"].(int64) < newMaps[j]["id"].(int64)})
	}
}

func makePointers(rows_len int) ([]any, []any) {
	content  := make([]any, rows_len)
 	pointers := make([]any, rows_len)
	for i := range content {
		pointers[i] = &content[i]
	}
	return content, pointers

}
