package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"Firstname":"Hendrik","Middlename":"Rizal","Lastname":"Array", "Age": 20, "Married" : false}`
	jsonBtyes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBtyes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
}
