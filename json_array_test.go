package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		Firstname:  "Hendrik",
		Middlename: "Array",
		Lastname:   "Rizal",
		Hobbies:    []string{"makan", "tidur"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
