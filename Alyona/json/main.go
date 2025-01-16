package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}
type Person struct {
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	IsEmployed   bool     `json:"isEmployed"`
	Address      Address  `json:"address"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

func main() {
	// Пример JSON строки
	jsonData := `{
					"name": "John",
					"age": 30,
					"isEmployed": true,
					"address": {
						"street": "123 Main St",
						"city": "New York"
					},
					"phoneNumbers": ["123-456-7890", "987-654-3210"]
					}`

	// Декодирование JSON в структуру
	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatalf("Error decoding JSON: %s", err)
	}

	fmt.Println("Decoded struct:", person)

	// Кодирование структуры обратно в JSON
	encodedJSON, err := json.Marshal(person)

	if err != nil {
		log.Fatalf("Error encoding struct to JSON: %s", err)
	}

	fmt.Println("Encoded JSON:", string(encodedJSON))
}
