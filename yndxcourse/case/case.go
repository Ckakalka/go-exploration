package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(1995, 12, 1, 0, 0, 0, 0, time.UTC)
	var greeting string
	switch year := date.Year(); {
	case year <= 1945:
		greeting = "Привет, старость!"
	case year <= 1964:
		greeting = "Привет, бумер!"
	case year <= 1980:
		greeting = "Привет, представитель X!"
	case year <= 1996:
		greeting = "Привет, миллениал!"
	case year <= 2012:
		greeting = "Привет, зумер!"
	case year > 2012:
		greeting = "Привет, альфа!"
	default:
		greeting = "Привет, неопознанный!"
	}
	fmt.Println(greeting)
}
