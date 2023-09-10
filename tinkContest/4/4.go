package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscanln(in, &n, &m)
	banknotes := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &banknotes[i])
	}
	stolenBanknotes := make([]int, 0, 10*2)
	for i := 1; i < int(math.Pow(3, float64(m))); i++ {
		stolenBanknotes = stolenBanknotes[:0]
		value := i
		stolenSum := 0
		for idx := 0; value != 0; idx, value = idx+1, value/3 {
			remainder := value % 3
			stolenSum += remainder * banknotes[idx]
			for j := 0; j < remainder; j++ {
				stolenBanknotes = append(stolenBanknotes, banknotes[idx])
			}
		}
		if stolenSum == n {
			fmt.Fprintln(out, len(stolenBanknotes))
			fmt.Fprint(out, stolenBanknotes[0])
			for j := 1; j < len(stolenBanknotes); j++ {
				fmt.Fprintf(out, " %d", stolenBanknotes[j])
			}
			fmt.Fprintln(out)
			return
		}
	}
	fmt.Fprintln(out, -1)
}
