package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	prev := 1
	curr := 1
	for i := 2; i <= n; i++ {
		temp := curr
		curr += prev
		prev = temp
	}
	return curr
}

func main() {
	var n int
	n = 2
	fmt.Println(climbStairs(n))

	n = 3
	fmt.Println(climbStairs(n))
}
