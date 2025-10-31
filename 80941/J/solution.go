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
	b := make([]int, n)

	sumA := 0
	for i := range n {
		fmt.Fscan(in, &a[i])
		sumA += a[i]
	}
	sumB := 0
	for i := range n {
		fmt.Fscan(in, &b[i])
		sumB += b[i]
	}

	if sumA > sumB {
		fmt.Fprintln(out, -1)
		return
	}

	answer := -1
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if check(&a, &b, n, mid) {
			answer = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	fmt.Fprintln(out, answer)
}

func check(a, b *[]int, n, k int) bool {
	left := 0
	extra := make([]int, n)

	for i := range n {
		start := max(0, i-k)
		end := min(n-1, i+k)

		currentNeed := (*a)[i]

		for left <= end && currentNeed > 0 {
			if left < start {
				left = start
				continue
			}
			if extra[left] < (*b)[left] {
				take := min(currentNeed, (*b)[left]-extra[left])
				currentNeed -= take
				extra[left] += take
			}
			if extra[left] >= (*b)[left] {
				left++
			}
		}

		if currentNeed > 0 {
			return false
		}
	}
	return true
}
