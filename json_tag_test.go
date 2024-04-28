package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Macbook pro",
		ImageUrl: "http://example.com",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Macbook pro","image_url":"http://example.com"}`
	jsonByte := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonByte, product)
	fmt.Println(product)
}
