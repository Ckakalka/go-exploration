package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	person := Person{
		Name:  "Алекс",
		Email: "alex@yandex.ru",
	}
	jsonPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(string(jsonPerson))
	}
}
