package psql

import (
	// "errors"
	"testing"
)

func TestDbConnection(t *testing.T) {
	tests := []struct {
		description string
		wantErr 	error
	}{
		{
		description: "Establishing Connection",
		wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := Psql_connect()
			if err != nil {
				t.Errorf("failed Psql_connect, got error: %v", err)
			}
		})
	}
}

func TestMapHandler(t *testing.T) {
	tests := []struct {
		description string
		mapKeys 	[]string
		slices 		[][]any
		want 		[]map[string]any
	}{
		{
			description: "convert slice of slices, with type 'any' into map",
			mapKeys: []string{"s1"},
			slices:	 [][]any{{15}, {"qwe"}},
			want: 	 []map[string]any{{"s1": 15}, {"s1": "qwe"}},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := convetIntoMap(test.slices, test.mapKeys)
			if result[0]["s1"] != test.want[0]["s1"] {
				t.Errorf("converting slices into map failed 'convertIntoMap'")
			}
		})
	}
}
