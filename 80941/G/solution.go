package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var m int
	fmt.Fscan(in, &m)
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(in, &b[i])
	}

	result := calc(n, &a, m, &b)
	fmt.Println(result)
}

func calc(n int, a *[]int, m int, b *[]int) int {
	aSorted := make([]int, n)
	copy(aSorted, *a)
	sort.Ints(aSorted)

	bSorted := make([]int, m)
	copy(bSorted, *b)
	sort.Ints(bSorted)

	prefixA := make([]int, n+1)
	for i := range n {
		prefixA[i+1] = prefixA[i] + aSorted[i]
	}

	prefixB := make([]int, m+1)
	for i := range m {
		prefixB[i+1] = prefixB[i] + bSorted[i]
	}

	var sum1 int
	for i := range n {
		sum1 += int(i+1) * sumOfAbs((*a)[i], &bSorted, prefixB)
	}

	var sum2 int
	for j := range m {
		sum2 += int(j+1) * sumOfAbs((*b)[j], &aSorted, prefixA)
	}

	return sum1 - sum2
}

func sumOfAbs(x int, arr *[]int, prefix []int) int {
	n := len(*arr)

	// Находим индекс первого элемента >= x с помощью бинарного поиска
	idx := sort.Search(n, func(i int) bool {
		return (*arr)[i] >= x
	})

	return x*idx - prefix[idx] + (prefix[n] - prefix[idx]) - x*(n-idx)
}
