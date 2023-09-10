package main

import (
	"bufio"
	"fmt"
	"os"
)

type Ghost struct {
	countGangs  int
	currentGang Gang
}

type Gang map[int]struct{}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	ghosts := make([]Ghost, 0, n)
	for i := 0; i < n; i++ {
		ghosts = append(ghosts, Ghost{1, make(Gang)})
		ghosts[i].currentGang[i] = struct{}{}
	}
	for i := 0; i < m; i++ {
		var q, x, y int
		fmt.Fscan(in, &q, &x)
		x--
		switch q {
		case 1:
			fmt.Fscan(in, &y)
			y--
			if _, ok := ghosts[x].currentGang[y]; ok {
				continue
			}
			for ghostIdx := range ghosts[y].currentGang {
				ghosts[x].currentGang[ghostIdx] = struct{}{}
			}
			for ghostIdx := range ghosts[x].currentGang {
				ghosts[ghostIdx].countGangs++
				ghosts[ghostIdx].currentGang = ghosts[x].currentGang
			}
		case 2:
			fmt.Fscan(in, &y)
			y--
			if _, ok := ghosts[x].currentGang[y]; ok {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		case 3:
			fmt.Fprintln(out, ghosts[x].countGangs)
		}
	}
}
