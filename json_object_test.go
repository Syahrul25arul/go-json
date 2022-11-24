package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	Firstname  string
	Middlename string
	Lastname   string
	Age        int
	Married    bool
	Hobbies    []string
	Address    []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Firstname:  "Rizal",
		Middlename: "hendrik",
		Lastname:   "array",
		Age:        50,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"Firstname":"Hendrik","Middlename":"Rizal","Lastname":"Array", "Age": 20, "Married" : false,"hobbies": ["makan","tidur"]}`
	jsonBtyes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBtyes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		Firstname: "Hendrik",
		Address: []Address{
			{
				Street:     "Jalan Biak",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan APO",
				Country:    "Indonesia",
				PostalCode: "1111",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"Firstname":"Hendrik","Middlename":"Rizal","Lastname":"Array", "Age": 20, "Married" : false,"hobbies": ["makan","tidur"], "Address":[{"Street":"Jalan Biak","Country":"Indonesia","PostalCode":"1111"}]}`
	jsonBtyes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBtyes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Address)
}

func TestJsonOnlyArrayComplex(t *testing.T) {
	Addreses := []Address{
		{
			Street:     "Jalan Biak",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan APO",
			Country:    "Indonesia",
			PostalCode: "1111",
		},
	}

	bytes, _ := json.Marshal(Addreses)
	fmt.Println(string(bytes))
}

func TestJsonOnlyArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Biak","Country":"Indonesia","PostalCode":"1111"}]`
	jsonBtyes := []byte(jsonString)

	Addreses := &[]Address{}
	err := json.Unmarshal(jsonBtyes, Addreses)

	if err != nil {
		panic(err)
	}

	fmt.Println(Addreses)
}
