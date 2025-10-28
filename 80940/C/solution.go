package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Interval struct {
	start, end, weight float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	intervals := make([]Interval, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &intervals[i].start, &intervals[i].end, &intervals[i].weight)
	}

	if n == 0 {
		fmt.Fprintf(out, "0.0000\n")
		return
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].end == intervals[j].end {
			return intervals[i].start < intervals[j].start
		}
		return intervals[i].end < intervals[j].end
	})

	check := make([]float64, n)
	check[0] = intervals[0].weight

	for i := 1; i < n; i++ {
		best := -1
		left, right := 0, i-1
		for left <= right {
			mid := (left + right) / 2

			if intervals[mid].end <= intervals[i].start {
				best = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		weight := intervals[i].weight
		if best != -1 {
			weight += check[best]
		}

		if check[i-1] > weight {
			check[i] = check[i-1]
		} else {
			check[i] = weight
		}
	}

	fmt.Fprintf(out, "%.5f\n", check[n-1])
}
