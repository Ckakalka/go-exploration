package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Row []int

type Column struct {
	table Table
	data  []int
}

func NewColumn(table Table) *Column {
	column := new(Column)
	column.table = table
	column.data = make([]int, column.Len())
	return column
}

func (column Column) Fill(number int) {
	for i := 0; i < column.Len(); i++ {
		column.data[i] = column.table.data[i][number]
	}
}

func (column Column) Len() int {
	return column.table.n
}

func (column Column) Less(i, j int) bool {
	return column.data[i] < column.data[j]
}

func (column Column) Swap(i, j int) {
	column.data[i], column.data[j] = column.data[j], column.data[i]
	temp := column.table.data[i]
	column.table.data[i] = column.table.data[j]
	column.table.data[j] = temp
}

type Table struct {
	n    int
	m    int
	data []Row
}

func (table Table) SortByColumn(columnNumber int) {
	column := NewColumn(table)
	column.Fill(columnNumber)
	sort.Stable(column)
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
		table.data = make([]Row, table.n)
		for j := 0; j < n; j++ {
			row := make(Row, m)
			for k := 0; k < m; k++ {
				fmt.Fscan(in, &row[k])
			}
			table.data[j] = row
		}
		var k int
		fmt.Fscan(in, &k)
		for j := 0; j < k; j++ {
			var click int
			fmt.Fscan(in, &click)
			table.SortByColumn(click - 1)
		}
		table.Print(out)
		fmt.Fprintln(out)
	}
}
