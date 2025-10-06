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

	var s string
	fmt.Fscan(in, &s)

	n := len(s)

	freq := make(map[rune]int)
	for _, v := range s {
		freq[v]++
	}

	total := n * (n - 1) / 2

	for _, count := range freq {
		if count == 1 {
			continue
		}

		total -= count * (count - 1) / 2

	}

	fmt.Fprintln(out, total+1)
}
