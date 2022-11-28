package gojson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		Firstname:  "Hendrik",
		Middlename: "Rizal",
		Lastname:   "Array",
	}

	encoder.Encode(customer)
}
