package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	friends := make([][]int, n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		friends[x] = append(friends[x], y)
		friends[y] = append(friends[y], x)

	}
	for i := 0; i < n; i++ {
		mutualFriends := make(map[int]int)
		for _, id := range friends[i] {
			for _, mutualFriend := range friends[id] {
				if mutualFriend == i {
					continue
				}
				isFriend := false
				for _, friend := range friends[i] {
					if friend == mutualFriend {
						isFriend = true
						break
					}
				}
				if isFriend {
					continue
				}
				mutualFriends[mutualFriend] += 1
			}
		}
		possibleFriends := make([]int, 0, 25)
		maxCount := -1
		for id, count := range mutualFriends {
			if count == maxCount {
				possibleFriends = append(possibleFriends, id+1)
			}
			if count > maxCount {
				maxCount = count
				possibleFriends = possibleFriends[:0]
				possibleFriends = append(possibleFriends, id+1)
			}
		}
		sort.Ints(possibleFriends)
		if len(possibleFriends) == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		fmt.Fprint(out, possibleFriends[0])
		for j := 1; j < len(possibleFriends); j++ {
			fmt.Fprintf(out, " %d", possibleFriends[j])
		}
		fmt.Fprintln(out)
	}
}
