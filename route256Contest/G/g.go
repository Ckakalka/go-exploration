package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func GetАllCards() map[string]struct{} {
	values := []string{"2", "3", "4", "5", "6", "7", "8",
		"9", "T", "J", "Q", "K", "A"}
	suits := []string{"S", "C", "D", "H"}
	allCards := make(map[string]struct{}, 52)
	for _, value := range values {
		for _, suit := range suits {
			allCards[value+suit] = struct{}{}
		}
	}
	return allCards
}

type Combination struct {
	rank  int
	value int
}

func DefineValue(value byte) int {
	symb := string(value)
	switch symb {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}
	return -1
}

func DefineCombination(a, b, c string) (combination Combination) {
	if a[0] == b[0] && a[0] == c[0] {
		combination.rank = 2
		combination.value = DefineValue(a[0])
		return
	}
	if a[0] == b[0] || a[0] == c[0] {
		combination.rank = 1
		combination.value = DefineValue(a[0])
		return
	}
	if b[0] == c[0] {
		combination.rank = 1
		combination.value = DefineValue(b[0])
		return
	}
	combination.rank = 0
	combination.value = int(math.Max(
		math.Max(float64(DefineValue(a[0])), float64(DefineValue(b[0]))),
		float64(DefineValue(c[0]))))
	return
}

func (firstCombo Combination) Compare(secondCombo Combination) int {
	if firstCombo.rank > secondCombo.rank {
		return 1
	}
	if firstCombo.rank == secondCombo.rank {
		if firstCombo.value > secondCombo.value {
			return 1
		}
		if firstCombo.value == secondCombo.value {
			return 0
		}
		return -1
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		cards := make([][2]string, n)
		availableCards := GetАllCards()
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &cards[j][0], &cards[j][1])
			delete(availableCards, cards[j][0])
			delete(availableCards, cards[j][1])
		}
		winCards := make([]string, 0, 52)
		for card := range availableCards {
			combinations := make([]Combination, 0, n)
			maxCombination := Combination{0, 2}
			for j := 0; j < n; j++ {
				combination := DefineCombination(cards[j][0], cards[j][1], card)
				combinations = append(combinations, combination)
				if combination.Compare(maxCombination) == 1 {
					maxCombination = combination
				}
			}
			if combinations[0] == maxCombination {
				winCards = append(winCards, card)
			}
		}
		fmt.Fprintln(out, len(winCards))
		for _, card := range winCards {
			fmt.Fprintln(out, card)
		}
	}
}
