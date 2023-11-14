package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	j := 0
	found := false
	for i := 0; i < len(s); i++ {
		found = false
		for j < len(t) {
			if s[i] == t[j] {
				found = true
				j++
				break
			}
			j++
		}
	}
	return found
}

func main() {
	var s, t string
	s = "abc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))

	s = "axc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
