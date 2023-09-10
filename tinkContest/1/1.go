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
	var n, s int
	fmt.Fscanln(in, &n, &s)
	maxRevolverPrice := 0
	var revolverPrice int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &revolverPrice)
		if maxRevolverPrice < revolverPrice && revolverPrice <= s {
			maxRevolverPrice = revolverPrice
		}
	}
	fmt.Fprintln(out, maxRevolverPrice)
}
