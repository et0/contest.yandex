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

	var n, k int
	fmt.Fscan(in, &n, &k)

	var theme int
	freq := make(map[int]int, n)
	for range n {
		fmt.Fscan(in, &theme)
		freq[theme]++
	}

	output := make([]int, 0, k)
	for len(output) < k {
		for i, v := range freq {
			if len(output) == k {
				break
			}

			output = append(output, i)

			if v == 1 {
				delete(freq, i)
			} else {
				freq[i]--
			}
		}
	}

	for _, v := range output {
		fmt.Fprintf(out, "%d ", v)
	}
}
