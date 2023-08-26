package main

import (
	"bufio"
	"fmt"
	"os"
)

const dot = 46

type Row []byte

type Table struct {
	n               int
	m               int
	processedRegion map[byte]struct{}
	previousRegion  byte
	data            []Row
}

func (table Table) Print(out *bufio.Writer) {
	for i := 0; i < table.n; i++ {
		fmt.Fprint(out, table.data[i][0])
		for j := 1; j < table.m; j++ {
			fmt.Fprintf(out, " %d", table.data[i][j])
		}
		fmt.Fprintln(out)
	}
}

func (table Table) Validation() bool {
	for i := 0; i < table.n; i++ {
		for j := 0; j < table.m; j++ {
			if !table.validateRegion(i, j) {
				return false
			}
			table.previousRegion = dot
		}
	}
	return true
}

func (table Table) validateRegion(x, y int) bool {
	if !(x >= 0 && x < table.n &&
		y >= 0 && y < table.m) {
		return true
	}
	if table.data[x][y] == dot {
		return true
	}
	if table.previousRegion != dot &&
		table.previousRegion != table.data[x][y] {
		return true
	}
	if table.previousRegion == dot {
		table.previousRegion = table.data[x][y]
		_, ok := table.processedRegion[table.data[x][y]]
		if ok {
			return false
		}
		table.processedRegion[table.data[x][y]] = struct{}{}
	}
	table.data[x][y] = dot
	isValid := table.validateRegion(x, y-2) &&
		table.validateRegion(x-1, y-1) &&
		table.validateRegion(x-1, y+1) &&
		table.validateRegion(x, y+2) &&
		table.validateRegion(x+1, y+1) &&
		table.validateRegion(x+1, y-1)
	return isValid
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	var n, m int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		table := new(Table)
		table.n = n
		table.m = m
		table.processedRegion = make(map[byte]struct{})
		table.previousRegion = dot
		table.data = make([]Row, table.n)
		var line string
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &line)
			row := []byte(line)
			table.data[j] = row
		}
		message := "YES"
		if !table.Validation() {
			message = "NO"
		}
		fmt.Fprintln(out, message)
	}
}
