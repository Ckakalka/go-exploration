package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	k := 0
	for i <= j {
		for i <= j && nums[i] != val {
			i++
			k++
		}
		for i <= j && nums[j] == val {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)
	fmt.Println(nums, k)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	k = removeElement(nums, val)
	fmt.Println(nums, k)
}
