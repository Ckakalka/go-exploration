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
	energyConsumption := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &energyConsumption[i])
	}
	sort.Ints(energyConsumption)
	releaseTimes := make([]int64, n)
	var t, l, sum int64
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &t, &l)
		for j := 0; j < n; j++ {
			if t >= releaseTimes[j] {
				sum += int64(energyConsumption[j]) * l
				releaseTimes[j] = t + l
				break
			}
		}
	}
	fmt.Fprintln(out, sum)
}
