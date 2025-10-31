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

	var p, root int
	parents := make([]int, n)
	children := make([][]int, n)
	for i := range n {
		fmt.Fscan(in, &p)

		if p == 0 {
			root = i
			continue
		}
		parents[i] = p
		children[p-1] = append(children[p-1], i)
	}

	inTime := make([]int, n)
	outTime := make([]int, n)
	timer := 0
	scan(root, &timer, &children, &inTime, &outTime)

	var m int
	fmt.Fscan(in, &m)

	var a, b int
	for range m {
		fmt.Fscan(in, &a, &b)
		if inTime[a-1] <= inTime[b-1] && outTime[a-1] >= outTime[b-1] {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}

func scan(root int, timer *int, children *[][]int, inTime, outTime *[]int) {
	(*timer)++
	(*inTime)[root] = *timer

	for _, c := range (*children)[root] {
		scan(c, timer, children, inTime, outTime)
	}

	(*timer)++
	(*outTime)[root] = *timer
}
