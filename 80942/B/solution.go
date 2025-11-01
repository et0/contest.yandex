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

	b := make([]int, n)
	t := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &b[i], &t[i])
	}

	var m, q int
	fmt.Fscan(in, &m)
	for m > 0 {
		m--

		fmt.Fscan(in, &q)

		left, right := 0, n-1
		idx := 0
		for left <= right {
			mid := left + (right-left)/2
			if b[mid] < q {
				idx = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		fmt.Fprintln(out, q*t[idx])
	}

}
