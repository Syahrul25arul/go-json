package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Mac Book Pro",
		ImageURL: "http://example.com/image",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","Name":"Apple Mac Book","image_url":"http://example.com/"}`
	jsonByte := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonByte, product)
	fmt.Println(product)
}
