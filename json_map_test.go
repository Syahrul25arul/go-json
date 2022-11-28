package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"p001", "name": "Apple Mac Book", "price" : 200000}`
	jsonBytes := []byte(jsonString)

	var result map[string]any
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
}

func TestMapDecode(t *testing.T) {
	product := map[string]any{
		"id":    "p001",
		"name":  "Apple Mac Book",
		"price": 200000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
