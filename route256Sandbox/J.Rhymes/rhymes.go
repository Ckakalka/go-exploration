package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)
	var s string
	dict := make([]string, 0, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		dict = append(dict, s)
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &s)
		maxSuffix := 0
		rhyme := ""
		for _, word := range dict {
			if word == s {
				continue
			}
			suffix := 0
			for k, j := len(word)-1, len(s)-1; k >= 0 && j >= 0; k, j = k-1, j-1 {
				if word[k] != s[j] {
					break
				}
				suffix++
			}
			if maxSuffix < suffix {
				maxSuffix = suffix
				rhyme = word
			}
		}
		if rhyme == "" {
			if s == dict[0] {
				rhyme = dict[1]
			} else {
				rhyme = dict[0]
			}
		}
		fmt.Fprintln(out, rhyme)
	}
}
