package main

import (
	"bufio"
	"fmt"
	"os"
)

type NumberToSkills map[int]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	var n int
	var skills NumberToSkills
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		skills = make(NumberToSkills, n)
		for j := 1; j <= n; j++ {
			var value int
			fmt.Fscan(in, &value)
			skills[j] = value
		}
		first := 0
		for len(skills) != 0 {
			ok := false
			for !ok {
				first++
				_, ok = skills[first]
			}
			firstSkill := skills[first]
			second := 101
			diff := 1024
			delete(skills, first)
			for number, skillVal := range skills {
				tempDiff := skillVal - firstSkill
				if tempDiff < 0 {
					tempDiff = -tempDiff
				}
				if tempDiff < diff {
					diff = tempDiff
					second = number
				} else if tempDiff == diff && number < second {
					second = number
				}
			}
			fmt.Fprintln(out, first, second)
			delete(skills, second)
		}
		fmt.Fprintln(out)
	}
}
