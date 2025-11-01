package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Season struct {
	serials int
	another int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	seasons := make([]Season, n)
	for i := range n {
		fmt.Fscan(in, &seasons[i].serials)
	}

	countAnother := 0
	for i := range n {
		fmt.Fscan(in, &seasons[i].another)
		countAnother += seasons[i].another
	}

	sort.Slice(seasons, func(i, j int) bool {
		return seasons[i].serials < seasons[j].serials
	})

	cumulativeWeight := 0
	serialsInSeason := seasons[0].serials

	for i := range n {
		cumulativeWeight += seasons[i].another
		if cumulativeWeight*2 >= countAnother {
			serialsInSeason = seasons[i].serials
			break
		}
	}

	payment := 0
	for i := range n {
		payment += abs(serialsInSeason-seasons[i].serials) * seasons[i].another
	}

	fmt.Fprintln(out, serialsInSeason, payment)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
