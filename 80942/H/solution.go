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

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}

	moneys := make([]int, n+1)
	for j := range n {
		if j+a[j] < n {
			moneys[j+a[j]]--
		}
		if j+1 < n {
			moneys[j+1]++
		}
	}

	sum, count := 0, 0
	for i := range n {
		count += moneys[i]
		sum += count * a[i]
	}

	fmt.Fprintln(out, sum)
}
