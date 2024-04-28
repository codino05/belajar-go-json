package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {

	customer := Customer{
		FirstName:  "eko",
		MiddleName: "kurniawan",
		LastName:   "khanedy",
		Hobbies: []string{
			"Gamming", "coding", "Reading",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"eko","MiddleName":"kurniawan","LastName":"khanedy","Age":0,"Married":false,"Hobbies":["Gamming","coding","Reading"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName: "eko",
		Address: []Address{
			{
				Street:     "jalan belum ada",
				Country:    "inodnesia",
				PostalCode: "52152",
			},
			{
				Street:     "jalan buntu",
				Country:    "indonesia",
				PostalCode: "42333",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"eko","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Address":[{"Street":"jalan belum ada","Country":"inodnesia","PostalCode":"52152"},{"Street":"jalan buntu","Country":"indonesia","PostalCode":"42333"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Address)
}

func TestOnlyJSONArrayComplexEncode(t *testing.T) {

	addresses := []Address{
		{
			Street:     "jalan belum ada",
			Country:    "inodnesia",
			PostalCode: "52152",
		},
		{
			Street:     "jalan buntu",
			Country:    "indonesia",
			PostalCode: "42333",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"jalan belum ada","Country":"inodnesia","PostalCode":"52152"},{"Street":"jalan buntu","Country":"indonesia","PostalCode":"42333"}]`
	jsonBytes := []byte(jsonString)

	anddress := &[]Address{}
	err := json.Unmarshal(jsonBytes, anddress)
	if err != nil {
		panic(err)
	}

	fmt.Println(anddress)
}
