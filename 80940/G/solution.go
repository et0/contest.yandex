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

	data := make([]int, n+1)
	data[0] = 1

	for num := range n {
		for sum := n; sum >= num+1; sum-- {
			data[sum] += data[sum-(num+1)]
		}
	}

	fmt.Fprintln(out, data[n])
}
