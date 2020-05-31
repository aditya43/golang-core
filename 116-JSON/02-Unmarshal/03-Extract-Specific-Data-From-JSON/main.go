package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	var p person

	jsonStr := `{"name":"Aditya","city":"Pune","favourite_bands":["Metallica","Lamb Of God","Manowar"],"phone_number":1234567890,"websites":[{"name":"AdiInviter Pro","url":"http://www.adiinviter.com"},{"name":"Github","url":"https://github.com/aditya43"}]}`

	err := json.Unmarshal([]byte(jsonStr), &p)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p.Name)
	fmt.Println(p.City)
}
