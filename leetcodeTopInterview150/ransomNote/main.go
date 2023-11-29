package main

import "fmt"

func getLetters(str string) map[byte]int {
	letters := make(map[byte]int)
	for i := range str {
		if _, ok := letters[str[i]]; !ok {
			letters[str[i]] = 0
		}
		letters[str[i]] += 1
	}
	return letters
}

func canConstruct(ransomNote string, magazine string) bool {
	lettersMagazine := getLetters(magazine)
	lettersRansomNote := getLetters(ransomNote)
	for letter, countRansomNote := range lettersRansomNote {
		if countMagazine := lettersMagazine[letter]; countMagazine < countRansomNote {
			return false
		}
	}
	return true
}

func main() {
	var ransomNote, magazine string
	ransomNote = "a"
	magazine = "b"
	fmt.Println(canConstruct(ransomNote, magazine))

	ransomNote = "aa"
	magazine = "ab"
	fmt.Println(canConstruct(ransomNote, magazine))

	ransomNote = "aa"
	magazine = "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}
