package main

import "fmt"

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	isLetter := func(s byte) bool {
		return ('a' <= s && s <= 'z') || ('A' <= s && s <= 'Z')
	}
	isAlpha := func(s byte) bool {
		return isLetter(s) || ('0' <= s && s <= '9')
	}
	var validDiff byte = 'a' - 'A'
	for left < right {
		for left < right && !isAlpha(s[left]) {
			left++
		}
		for left < right && !isAlpha(s[right]) {
			right--
		}
		var diff byte
		if s[left] > s[right] {
			diff = s[left] - s[right]
		} else {
			diff = s[right] - s[left]
		}

		if !(diff == 0 || (isLetter(s[left]) && diff == validDiff)) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))

	s = "race a car"
	fmt.Println(isPalindrome(s))

	s = "0P"
	fmt.Println(isPalindrome(s))
}
