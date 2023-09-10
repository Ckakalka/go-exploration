package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsEqualMap(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscanln(in, &n)
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &cards[i])
	}
	winCards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &winCards[i])
	}
	var leftIdx int
	for i := 0; i < n; i++ {
		if cards[i] != winCards[i] {
			leftIdx = i
			break
		}
	}
	var rightIdx int
	for i := n - 1; i >= 0; i-- {
		if cards[i] != winCards[i] {
			rightIdx = i
			break
		}
	}
	winNonDecreasing := true
	cardsToCount := make(map[int]int)
	winCardsToCount := make(map[int]int)
	previousCard := 0
	for i := leftIdx; i <= rightIdx; i++ {
		if !(previousCard <= winCards[i]) {
			winNonDecreasing = false
		}
		previousCard = winCards[i]

		if _, ok := cardsToCount[cards[i]]; !ok {
			cardsToCount[cards[i]] = 0
		}
		cardsToCount[cards[i]] += 1
		if _, ok := winCardsToCount[winCards[i]]; !ok {
			winCardsToCount[winCards[i]] = 0
		}
		winCardsToCount[winCards[i]] += 1
	}
	if !winNonDecreasing {
		fmt.Fprintln(out, "NO")
		return
	}
	if IsEqualMap(cardsToCount, winCardsToCount) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
