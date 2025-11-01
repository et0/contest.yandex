package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	freq  int
	bumps int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}

	freq := make([]int, n+2)
	for range m {
		var l, r int
		fmt.Fscan(in, &l, &r)
		freq[l]++
		if r+1 <= n {
			freq[r+1]--
		}
	}

	for i := 1; i <= n; i++ {
		freq[i] += freq[i-1]
	}

	var totalD int
	for i := 1; i <= n; i++ {
		totalD += a[i-1] * freq[i]
	}

	pairs := make([]Pair, n)
	for i := range n {
		pairs[i] = Pair{freq[i+1], a[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	for i := 0; i < n && k > 0; i++ {
		if pairs[i].freq == 0 {
			break
		}
		repair := min(pairs[i].bumps, k)
		totalD -= repair * pairs[i].freq
		k -= repair
	}

	fmt.Fprintln(out, totalD)
}
