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

	table := make([][]int, n)
	memo := make([][]int, n)
	for i := range n {
		table[i] = make([]int, m)
		memo[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(in, &table[i][j])
		}
	}

	dirs := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	maxLen := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}

		current := table[i][j]
		maxPath := 1

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < n && nj >= 0 && nj < m && table[ni][nj] == current+1 {
				pathLen := 1 + dfs(ni, nj)
				if pathLen > maxPath {
					maxPath = pathLen
				}
			}
		}

		memo[i][j] = maxPath
		return maxPath
	}

	for i := range n {
		for j := range m {
			length := dfs(i, j)
			if length > maxLen {
				maxLen = length
			}
		}
	}

	fmt.Fprintln(out, maxLen)
}
