package main

import "fmt"

func RemoveDeduplicates(input []string) []string {
	mapInput := make(map[string]struct{}, len(input))
	result := make([]string, 0)
	for _, v := range input {
		if _, ok := mapInput[v]; !ok {
			result = append(result, v)
			mapInput[v] = struct{}{}
		}
	}
	return result
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	output := RemoveDeduplicates(input)
	fmt.Println(output)
}
