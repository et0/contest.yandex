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

	var n, m int
	fmt.Fscan(in, &n, &m)

	// сумма в строке и столбце
	rowSum := make([]int, n)
	colSum := make([]int, m)

	// кол-во вопросов в строке и столбце
	rowQ := make([]int, n)
	colQ := make([]int, m)

	grid := make([]string, n)
	for y := 0; y < n; y++ {
		fmt.Fscan(in, &grid[y])

		for x, v := range grid[y] {
			switch v {
			case '+':
				colSum[x]++
				rowSum[y]++
			case '-':
				colSum[x]--
				rowSum[y]--
			case '?':
				rowQ[y]++
				colQ[x]++
			}
		}
	}

	resultDiff := -1000
	for y := range n {
		for x := range m {
			diff := (rowSum[y] + rowQ[y]) - (colSum[x] - colQ[x])

			if grid[y][x] == '?' {
				option1 := (rowSum[y] + (rowQ[y] - 1) - 1) - (colSum[x] - colQ[x]) // ? = -
				option2 := (rowSum[y] + rowQ[y]) - (colSum[x] - (colQ[x] - 1) + 1) // ? = +

				diff = max(option1, option2)
			}

			if diff > resultDiff {
				resultDiff = diff
			}
		}
	}

	fmt.Fprintln(out, resultDiff)

}
