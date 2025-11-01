package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type Delta struct {
	dx, dy int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, d int
	fmt.Fscan(in, &n, &d)

	points := make([]Point, n)
	for i := range n {
		fmt.Fscan(in, &points[i].x, &points[i].y)
	}

	deltas := make([]Delta, 0, 1000)
	for dx := 0; dx*dx <= d; dx++ {
		rem := d - dx*dx
		dy := 0
		for dy*dy < rem {
			dy++
		}

		if dy*dy == rem {
			if dx == 0 && dy == 0 {
				//
			} else if dx == 0 {
				deltas = append(deltas, Delta{0, dy}, Delta{0, -dy})
			} else if dy == 0 {
				deltas = append(deltas, Delta{dx, 0}, Delta{-dx, 0})
			} else {
				deltas = append(deltas, Delta{dx, dy}, Delta{dx, -dy}, Delta{-dx, dy}, Delta{-dx, -dy})
			}
		}
	}

	pointCount := make(map[Point]int)
	result := 0

	for _, p := range points {
		for _, delta := range deltas {
			result += pointCount[Point{p.x + delta.dx, p.y + delta.dy}]
		}
		pointCount[p]++
	}

	fmt.Fprintln(out, result)
}
