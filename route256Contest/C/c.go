package main

import (
	"bufio"
	"fmt"
	"os"
)

type NumberToSkills map[int]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	var n int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		minT := 15
		maxT := 30
		var s string
		var val int
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &s, &val)
			if s == ">=" {
				if val > minT {
					minT = val
				}
			} else {
				if val < maxT {
					maxT = val
				}
			}
			if minT > maxT {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, minT)
			}
		}
		fmt.Fprintln(out)
	}
}
