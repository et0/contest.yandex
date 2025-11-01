package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	var x float64
	fmt.Fscan(in, &n, &m, &x)

	intervals := make([]struct{ start, end float64 }, n)

	for i := range n {
		var a, b, v float64
		fmt.Fscan(in, &a, &b, &v)

		if a > b {
			a, b = b, a
			v = -v
		}

		end := math.Max((x-b)/v, (x-a)/v)
		if end < 0 {
			continue
		}

		start := math.Min((x-b)/v, (x-a)/v)
		if start < 0 {
			start = 0
		}

		intervals[i] = struct{ start, end float64 }{start, end}
	}

	if len(intervals) > 0 {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].start < intervals[j].start
		})

		merged := make([]struct{ start, end float64 }, 0, len(intervals))
		merged = append(merged, intervals[0])

		for i := 1; i < len(intervals); i++ {
			last := &merged[len(merged)-1]
			if intervals[i].start <= last.end {
				if intervals[i].end > last.end {
					last.end = intervals[i].end
				}
			} else {
				merged = append(merged, intervals[i])
			}
		}
		intervals = merged
	}

	var t float64
	for range m {
		fmt.Fscan(in, &t)

		i := sort.Search(len(intervals), func(j int) bool {
			return intervals[j].end >= t
		})

		if i < len(intervals) && intervals[i].start <= t {
			fmt.Fprintf(out, "%.7f\n", intervals[i].end)

			continue
		}

		fmt.Fprintf(out, "%.7f\n", t)
	}
}
