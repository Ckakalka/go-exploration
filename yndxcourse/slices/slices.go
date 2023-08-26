package main

import "fmt"

func main() {
	slice := make([]int, 100, 100)
	for i := range slice {
		slice[i] = i + 1
	}
	fmt.Println(slice)
	slice = append(slice[:10], slice[90:]...)
	mid := len(slice) / 2
	for i := 0; i < mid; i++ {
		j := len(slice) - 1 - i
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println(slice)
}
