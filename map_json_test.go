package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro", "price": 20000000}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJSONMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Apple mac book pro",
		"price": 20000000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
