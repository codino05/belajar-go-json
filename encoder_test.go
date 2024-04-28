package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStramEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "eko",
		MiddleName: "kurniawan",
		LastName:   "khanedy",
	}
	encoder.Encode(customer)
}
