package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x1, y1, x2, y2 float64
	fmt.Fscan(in, &x1, &y1, &x2, &y2)

	dx := int(math.Abs(x1 - x2))
	dy := int(math.Abs(y1 - y2))

	result := 0
	if x1 == x2 && y1 == y2 {

	} else if x1 != x2 && y1 != y2 {
		result = 3*((dx-1)+(dy-1)) + 1
	} else if dx == 0 {
		// Вертикальное движение
		result = 3 * (dy - 1)
	} else if dy == 0 {
		// Горизонтальное движение
		result = 3 * (dx - 1)
	}

	fmt.Fprintln(out, result)
}
