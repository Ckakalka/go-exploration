package main

import "fmt"

func maxProfit(prices []int) int {
	maxDiff := 0
	cumDiff := 0
	for i := len(prices) - 1; i > 0; i-- {
		if cumDiff+(prices[i]-prices[i-1]) > 0 {
			cumDiff += (prices[i] - prices[i-1])
		} else {
			cumDiff = 0
		}
		if cumDiff > maxDiff {
			maxDiff = cumDiff
		}
	}
	return maxDiff
}

func main() {
	nums1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums1))

	nums2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(nums2))
}
