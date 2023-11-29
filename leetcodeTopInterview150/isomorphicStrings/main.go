package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	lastPositionS := make(map[byte]int)
	lastPositionT := make(map[byte]int)
	for i := range s {
		posS, okS := lastPositionS[s[i]]
		posT, okT := lastPositionT[t[i]]
		if okS != okT {
			return false
		}
		if posS != posT {
			return false
		}
		lastPositionS[s[i]] = i
		lastPositionT[t[i]] = i
	}
	return true
}

func main() {
	var s, t string
	s = "egg"
	t = "add"
	fmt.Println(isIsomorphic(s, t))

	s = "foo"
	t = "bar"
	fmt.Println(isIsomorphic(s, t))

	s = "paper"
	t = "title"
	fmt.Println(isIsomorphic(s, t))
}
