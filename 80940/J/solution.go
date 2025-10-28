package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Shop struct {
	P, R, Q, F int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, L int
	fmt.Fscan(in, &N, &L)

	shops := make([]Shop, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &shops[i].P, &shops[i].R, &shops[i].Q, &shops[i].F)
	}

	maxMeters := L + 100

	check := make([][]int, N+1)
	purchase := make([][][]int, N+1)
	for i := range check {
		check[i] = make([]int, maxMeters+1)
		purchase[i] = make([][]int, maxMeters+1)
		for j := range check[i] {
			check[i][j] = math.MaxInt32
			purchase[i][j] = make([]int, N)
		}
	}

	check[0][0] = 0

	for i := 1; i <= N; i++ {
		shop := shops[i-1]

		for m := range maxMeters {
			if check[i-1][m] < check[i][m] {
				check[i][m] = check[i-1][m]
				copy(purchase[i][m], purchase[i-1][m])
			}
		}

		for prevMeters := range maxMeters {
			if check[i-1][prevMeters] == math.MaxInt32 {
				continue
			}

			for k := 1; k <= shop.F; k++ {
				newMeters := prevMeters + k
				if newMeters > maxMeters {
					continue
				}

				var cost int
				if k < shop.R {
					cost = k * shop.P
				} else {
					cost = k * shop.Q
				}

				totalCost := check[i-1][prevMeters] + cost
				if totalCost < check[i][newMeters] {
					check[i][newMeters] = totalCost
					copy(purchase[i][newMeters], purchase[i-1][prevMeters])
					purchase[i][newMeters][i-1] = k
				}
			}
		}
	}

	minCost := math.MaxInt32
	bestMeters := -1
	for m := L; m <= maxMeters; m++ {
		if check[N][m] < minCost {
			minCost = check[N][m]
			bestMeters = m
		}
	}

	if minCost == math.MaxInt32 {
		fmt.Fprintln(out, -1)
		return
	}

	fmt.Fprintln(out, minCost)
	for i := range N {
		fmt.Fprint(out, purchase[N][bestMeters][i])
		if i < N-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}
