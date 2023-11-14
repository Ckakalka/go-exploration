package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	subArr := make([]int, k)
	for i := 0; i < k; i++ {
		subArr[i] = nums[len(nums)-k+i]
	}
	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = subArr[i]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	rotate(nums1, k1)
	fmt.Println(nums1)

	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(nums2, k2)
	fmt.Println(nums2)
}
