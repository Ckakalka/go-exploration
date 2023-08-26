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
	var a, b int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
