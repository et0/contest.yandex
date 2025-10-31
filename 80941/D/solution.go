package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type table struct {
	value float64
	index int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var p float64
	fmt.Fscan(in, &n, &p)

	var value float64
	tables := make([]table, n)
	for i := range n {
		fmt.Fscan(in, &value)
		tables[i] = table{value, i}
	}

	sort.Slice(tables, func(i, j int) bool {
		return tables[i].value < tables[j].value
	})

	bestDiff := math.MaxFloat64
	findI, findJ := -1, -1
	for j := range n {
		cj := tables[j].value

		target := p * cj
		pos := sort.Search(n, func(i int) bool {
			return tables[i].value >= target
		})

		try := make([]int, 0, 3)
		if pos < n {
			try = append(try, pos)
		}
		if pos > 0 {
			try = append(try, pos-1)
		}
		if pos < n-1 {
			try = append(try, pos+1)
		}

		for _, i := range try {
			if i == j {
				continue
			}

			diff := math.Abs(tables[i].value/cj - p)
			if diff >= bestDiff {
				continue
			}

			bestDiff = diff
			findI, findJ = tables[i].index+1, tables[j].index+1
		}
	}

	fmt.Fprintln(out, findI, findJ)
}
