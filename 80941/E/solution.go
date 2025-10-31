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

	var p int
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &p)
		children[p] = append(children[p], i)
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var result int
	check(0, &result, &children, &a)

	fmt.Println(result)
}

func check(node int, result *int, children *[][]int, a *[]int) int {
	if len((*children)[node]) == 0 {
		*result += abs((*a)[node])

		return (*a)[node]
	}

	var sumChildren int
	for _, child := range (*children)[node] {
		sumChildren += check(child, result, children, a)
	}

	diff := (*a)[node] - sumChildren
	*result += abs(diff)

	return (*a)[node]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
