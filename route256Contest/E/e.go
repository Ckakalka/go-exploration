package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Friend struct {
	a   int
	idx int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	friends := make([]Friend, n)
	for i := 0; i < n; i++ {
		friend := Friend{}
		fmt.Fscan(in, &friend.a)
		friend.idx = i
		friends[i] = friend
	}
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].a < friends[j].a
	})
	minCard := 1
	maxCard := m
	cards := make([]int, n)
	ok := true
	for _, friend := range friends {
		if minCard <= friend.a {
			cards[friend.idx] = friend.a + 1
			minCard = friend.a + 2
		} else {
			cards[friend.idx] = minCard
			minCard++
		}
		if cards[friend.idx] > maxCard {
			ok = false
			break
		}
	}
	if ok {
		fmt.Fprint(out, cards[0])
		for i := 1; i < n; i++ {
			fmt.Fprintf(out, " %d", cards[i])
		}
	} else {
		fmt.Fprint(out, -1)
	}
	fmt.Fprintln(out)
}
