package main

import "fmt"

func majorityElement(nums []int) int {
	numsToCount := make(map[int]int)
	maxNum := -1
	for _, num := range nums {
		if _, ok := numsToCount[num]; !ok {
			numsToCount[num] = 0
		}
		numsToCount[num] += 1
		if num != maxNum && numsToCount[maxNum] < numsToCount[num] {
			maxNum = num
		}
	}
	return maxNum
}

func main() {
	nums1 := []int{3, 2, 3}
	fmt.Println(majorityElement(nums1))

	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums2))
}
