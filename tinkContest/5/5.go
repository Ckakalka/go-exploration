package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Road struct {
	to     int
	length int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscanln(in, &n, &m)
	roads := make([][]Road, n)
	for i := range roads {
		roads[i] = make([]Road, 0)
	}
	for i := 0; i < m; i++ {
		var v, u, w int
		fmt.Fscan(in, &v, &u, &w)
		v--
		u--
		roads[v] = append(roads[v], Road{u, w})
		roads[u] = append(roads[u], Road{v, w})
	}
	citiesInsideStates := make([]bool, n)
	states := make([][]int, 0, n)
	var dfs func(v int)
	dfs = func(v int) {
		idxState := len(states) - 1
		states[idxState] = append(states[idxState], v)
		citiesInsideStates[v] = true
		for _, road := range roads[v] {
			if !citiesInsideStates[road.to] {
				dfs(road.to)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !citiesInsideStates[i] {
			states = append(states, make([]int, 0))
			dfs(i)
		}
	}
	minSpanningRoadLength := math.MaxInt
	for _, state := range states {
		maxSpanningTree := make(map[int]struct{})
		maxSpanningTree[state[0]] = struct{}{}
		for len(maxSpanningTree) != len(state) {
			maxRoad := Road{-1, 0}
			for v := range maxSpanningTree {
				for _, road := range roads[v] {
					if _, ok := maxSpanningTree[road.to]; ok {
						continue
					}
					if road.length > maxRoad.length {
						maxRoad = road
					}
				}
			}
			maxSpanningTree[maxRoad.to] = struct{}{}
			if maxRoad.length < minSpanningRoadLength {
				minSpanningRoadLength = maxRoad.length
			}
		}
	}
	fmt.Fprintln(out, minSpanningRoadLength-1)
}
