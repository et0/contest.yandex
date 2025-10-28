package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	grid := make([]string, n)
	for i := range n {
		fmt.Fscan(in, &grid[i])
	}

	if grid[0][0] == 'W' && grid[0][1] == 'W' && grid[0][2] == 'W' {
		fmt.Fprintln(out, 0)
		return
	}

	way := [3]int{}
	for i, v := range grid[0] {
		switch v {
		case '.':
			way[i] = 0
		case 'C':
			way[i] = 1
		case 'W':
			way[i] = -1
		}
	}

	for n := 1; n < len(grid); n++ {
		nextWay := [3]int{}

		for m, w := range grid[n] {
			if w == 'W' {
				nextWay[m] = -1
				continue
			}

			switch m {
			case 0:
				nextWay[m] = max(way[0], way[1])
			case 1:
				nextWay[m] = max(way[0], way[1], way[2])
			case 2:
				nextWay[m] = max(way[1], way[2])
			}

			if w == 'C' && nextWay[m] != -1 {
				nextWay[m] += 1
			}
		}

		if nextWay[0] == -1 && nextWay[1] == -1 && nextWay[2] == -1 {
			break
		}

		way[0], way[1], way[2] = nextWay[0], nextWay[1], nextWay[2]
	}

	fmt.Fprintln(out, max(way[0], way[1], way[2]))
}
