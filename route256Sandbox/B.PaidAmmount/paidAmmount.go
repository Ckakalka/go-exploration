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
	var t int
	fmt.Fscan(in, &t)
	var n int
	var products PriceToCount
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		products = make(PriceToCount, n)
		for j := 0; j < n; j++ {
			var value int
			fmt.Fscan(in, &value)
			_, ok := products[value]
			if !ok {
				products[value] = 0
			}
			products[value] += 1
		}
		totalPrice := 0
		for price, count := range products {
			totalPrice += (count - count/3) * price
		}
		fmt.Fprintln(out, totalPrice)
	}
}
