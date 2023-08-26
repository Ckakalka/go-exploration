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
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		tasks := make([]bool, n)
		var taskNumber int
		previousTaskNumber := -1
		message := "Yes"
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &taskNumber)
			if tasks[taskNumber-1] != false && previousTaskNumber != taskNumber {
				message = "NO"
			}
			tasks[taskNumber-1] = true
			previousTaskNumber = taskNumber
		}
		fmt.Fprintln(out, message)
	}
}
