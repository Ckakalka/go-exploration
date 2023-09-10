package main

import (
	"bufio"
	"fmt"
	"os"
)

type PriceToCount map[int]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var s_temp string
	fmt.Fscanln(in, &s_temp)
	s := []rune(s_temp)
	var n int
	fmt.Fscan(in, &n)
	var start, end int
	var r_temp string
	var r []rune
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &start, &end, &r_temp)
		r = []rune(r_temp)
		for j := start - 1; j < end; j++ {
			s[j] = r[j-start+1]
		}
	}
	fmt.Fprintln(out, string(s))
}
