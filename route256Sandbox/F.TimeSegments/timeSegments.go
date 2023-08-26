package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type TimeSegment struct {
	date        time.Time
	isBeginning bool
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		var line string
		message := "Yes"
		dates := make([]TimeSegment, 0, 2*n)
		expectedCountEqualDates := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &line)
			datesStr := strings.Split(line, "-")
			pairDates := make([]time.Time, len(datesStr))
			for k, dateStr := range datesStr {
				var err error
				pairDates[k], err = time.Parse("15:04:05", dateStr)
				if err != nil {
					message = "NO"
					break
				}
			}
			if pairDates[0].Equal(pairDates[1]) {
				expectedCountEqualDates++
			}
			if !(pairDates[0].Before(pairDates[1]) ||
				pairDates[0].Equal(pairDates[1])) {
				message = "NO"
			}
			dates = append(dates, TimeSegment{pairDates[0], true},
				TimeSegment{pairDates[1], false})
		}
		if message != "NO" {
			comparator := func(i, j int) bool {
				return dates[i].date.Before(dates[j].date)
			}
			sort.SliceStable(dates, comparator)
			expectedBeginning := true
			countEqualDates := 0
			for k := 0; k < len(dates)-1; k++ {
				if expectedBeginning != dates[k].isBeginning {
					message = "NO"
					break
				}
				if dates[k].date.Equal(dates[k+1].date) {
					countEqualDates++
				}
				expectedBeginning = !expectedBeginning
			}
			if expectedCountEqualDates != countEqualDates {
				message = "NO"
			}
		}
		fmt.Fprintln(out, message)
	}
}
