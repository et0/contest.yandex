package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	cols := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &cols[i])
	}

	deque := make([]int, 0, n)
	minTower := make([]int, 0, n-k+1)
	sumTower := make([]int, 0, n-k+1)
	sum := 0
	for i := 0; i < n; i++ {
		if len(deque) > 0 && deque[0] == i-k {
			deque = deque[1:]
		}

		for len(deque) > 0 && cols[deque[len(deque)-1]] > cols[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		sum += cols[i]

		if i >= k-1 {
			minTower = append(minTower, cols[deque[0]])
			sumTower = append(sumTower, sum)

			sum -= cols[i+1-k]
		}
	}

	dp := make([]int, n+1)
	prev := make([]int, n+1)

	for i := range dp {
		dp[i] = -1
		prev[i] = -1
	}
	dp[0] = 0

	for i := 0; i <= n; i++ {
		if dp[i] == -1 {
			continue
		}

		// строим башню
		if i <= n-k {
			newVal := dp[i] + minTower[i]*sumTower[i]
			if newVal > dp[i+k] {
				dp[i+k] = newVal
				prev[i+k] = i
			}
		}

		// или пропускаем
		if i < n && dp[i] > dp[i+1] {
			dp[i+1] = dp[i]
			prev[i+1] = i
		}
	}

	towers := make([]int, 0)
	pos := n

	for pos > 0 {
		if prev[pos] == -1 {
			pos--
			continue
		}

		if pos-prev[pos] == k {
			towers = append(towers, prev[pos]+1)
			pos = prev[pos]
		} else {
			pos--
		}
	}

	slices.Reverse(towers)

	fmt.Fprintln(out, len(towers))
	for i := 0; i < len(towers); i++ {
		fmt.Fprintf(out, "%d ", towers[i])
	}
	fmt.Fprintln(out)
}
