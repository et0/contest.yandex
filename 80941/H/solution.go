package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	total := 0
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
		total += a[i]
	}

	if n == 1 {
		fmt.Fprintln(out, 1)
		return
	}

	var u, v int
	conn := make([][]int, n)
	for range n - 1 {
		fmt.Fscan(in, &u, &v)
		conn[u-1] = append(conn[u-1], v-1)
		conn[v-1] = append(conn[v-1], u-1)
	}

	sub, parent := make([]int, n), make([]int, n)
	checkSub(0, -1, &sub, &parent, &conn, &a)

	firstMax := make([][2]int, n)
	for u := range n {
		max1, max2 := 0, 0
		for _, v := range conn[u] {
			if v == parent[u] {
				continue
			}
			if sub[v] > max1 {
				max2 = max1
				max1 = sub[v]
			} else if sub[v] > max2 {
				max2 = sub[v]
			}
		}
		firstMax[u] = [2]int{max1, max2}
	}

	var changeRoot func(int, int)
	secondMax, answer, maxValue := make([]int, n), -1, math.MaxInt64
	changeRoot = func(u, p int) {
		queue := max(firstMax[u][0], a[u])

		if u != 0 && secondMax[u] > queue {
			queue = secondMax[u]
		}

		if queue < maxValue {
			maxValue = queue
			answer = u
		}

		for _, v := range conn[u] {
			if v == p {
				continue
			}

			candidate := total - sub[v]

			if sub[v] == firstMax[u][0] {
				candidate = max(candidate, firstMax[u][1])
			} else {
				candidate = max(candidate, firstMax[u][0])
			}

			candidate = max(candidate, a[u])
			secondMax[v] = candidate

			changeRoot(v, u)
		}
	}
	changeRoot(0, -1)

	fmt.Fprintln(out, answer+1)
}

func checkSub(u, p int, sub, parent *[]int, conn *[][]int, a *[]int) {
	(*parent)[u] = p
	(*sub)[u] = (*a)[u]
	for _, v := range (*conn)[u] {
		if v == p {
			continue
		}
		checkSub(v, u, sub, parent, conn, a)
		(*sub)[u] += (*sub)[v]
	}
}
