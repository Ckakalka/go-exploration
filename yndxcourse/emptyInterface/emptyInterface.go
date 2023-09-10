package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch A := a.(type) {
	case int:
		return A * b
	case string:
		return strings.Repeat(A, b)
	default:
		return nil
	}
}

func main() {
	f := 5.5
	a := 3
	str := "ABC"
	fmt.Println(Mul(a, a))
	fmt.Println(Mul(str, a))
	fmt.Println(Mul(f, a))
}
