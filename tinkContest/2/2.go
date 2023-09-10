package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	keyword := "sheriff"
	keyRuneToCount := make(map[rune]int)
	runeToCount := make(map[rune]int)
	for _, symb := range keyword {
		if _, ok := keyRuneToCount[symb]; !ok {
			keyRuneToCount[symb] = 0
			runeToCount[symb] = 0
		}
		keyRuneToCount[symb] += 1
	}
	var s string
	fmt.Fscanln(in, &s)
	for _, symb := range s {
		if _, ok := runeToCount[symb]; ok {
			runeToCount[symb] += 1
		}
	}
	minMult := len(s)
	for symb, count := range runeToCount {
		mult := count / keyRuneToCount[symb]
		if mult < minMult {
			minMult = mult
		}
	}
	fmt.Fprintln(out, minMult)
}
