package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	prev := nums[0]
	i := 1
	j := i
	for j < len(nums) {
		if nums[i] <= prev {
			for j < len(nums) && nums[j] == prev {
				j++
			}
			if j < len(nums) {
				nums[i] = nums[j]
				prev = nums[i]
				i++
				j++
			}
		} else {
			prev = nums[i]
			i++
			j++
		}
	}
	return i
}

func main() {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k = removeDuplicates(nums)
	fmt.Println(nums, k)
}
