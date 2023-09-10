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
	var t int
	fmt.Fscan(in, &t)
	var k, n, m int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &k, &n, &m)
		visibleRelief := make([][]byte, n)
		for x := 0; x < n; x++ {
			visibleRelief[x] = make([]byte, m)
			for y := 0; y < m; y++ {
				visibleRelief[x][y] = '.'
			}
		}
		var line string
		for j := 0; j < k; j++ {
			relief := make([][]byte, n)
			for counter := 0; counter < n; counter++ {
				fmt.Fscan(in, &line)
				relief[counter] = []byte(line)
			}
			for x := 0; x < n; x++ {
				for y := 0; y < m; y++ {
					if visibleRelief[x][y] == '.' {
						visibleRelief[x][y] = relief[x][y]
					}
				}
			}
		}
		for x := 0; x < n; x++ {
			fmt.Fprintln(out, string(visibleRelief[x]))
		}
		fmt.Fprintln(out)
	}
}
