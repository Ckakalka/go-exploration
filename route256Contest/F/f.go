package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type CategoriesTrees []CategoriesTree

type CategoriesTree struct {
}

type NodesArr []Node

type Node struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IdParent int    `json:"parent,omitempty"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscanln(in, &t)
	var m int
	for i := 0; i < t; i++ {
		fmt.Fscanln(in, &m)
		fmt.Fprintf(out, "m = %d\n", m)
		var nodesArr NodesArr
		input := ""
		for j := 0; j < m; j++ {
			var line string
			fmt.Fscanln(in, &line)
			fmt.Fprintf(out, "line %d %s\n", j+1, line)
			input += line + "\n"
		}
		fmt.Fprintln(out, input)
		err := json.Unmarshal([]byte(input), &nodesArr)
		if err != nil {
			fmt.Fprintln(out, err)
		}
		fmt.Fprintln(out, nodesArr)
	}
}
