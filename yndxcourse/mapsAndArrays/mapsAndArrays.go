package main

import "fmt"

func getIdxPairs(data []int, k int) []int {
	mapData := make(map[int][]int)
	for i, v := range data {
		_, ok := mapData[v]
		if !ok {
			mapData[v] = make([]int, 0)
		}
		mapData[v] = append(mapData[v], i)
	}
	result := make([]int, 0)
	for key, vals := range mapData {
		tVals, ok := mapData[k-key]
		if ok {
			for _, v := range vals {
				for _, tV := range tVals {
					result = append(result, v, tV)
				}
			}
		}
		delete(mapData, key)
		delete(mapData, k-key)
	}
	return result
}

func main() {
	a := []int{1, 1, 1, 9, 2, 8, 3, 7}
	k := 10
	result := getIdxPairs(a[:], k)
	fmt.Println(result)
}
