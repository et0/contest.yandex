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

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + a[i-1]
	}

	total := prefix[n]
	minDiff := total
	bestLeft, bestRight := 1, n
	for k := 1; k <= n-1; k++ {
		target := total - prefix[k]

		low, high := 1, k
		for low <= high {
			mid := (low + high) / 2

			diff := abs(prefix[mid] + prefix[k] - total)
			if diff < minDiff || (diff == minDiff && mid < bestLeft) {
				minDiff = diff
				bestLeft, bestRight = mid, k+1
			}

			if prefix[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		for _, l := range []int{max(1, low-2), min(k, low+2)} {
			if l >= 1 && l <= k {
				diff := abs(prefix[l] + prefix[k] - total)
				if diff < minDiff || (diff == minDiff && l < bestLeft) {
					minDiff = diff
					bestLeft, bestRight = l, k+1
				}
			}
		}
	}

	fmt.Fprintln(out, minDiff, bestLeft, bestRight)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
