package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type word struct {
	w, h float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		n    int
		w, h float64
		a, b float64
	)
	fmt.Fscan(in, &n, &w, &h)

	words := make([]word, n)
	for i := range n {
		fmt.Fscan(in, &a, &b)

		words[i] = word{a, b}
	}

	var (
		left         float64 = 0.0000009
		right        float64 = 1000000000.0
		mid, prevMid float64
	)
	prevMid = left
	for range 100 {
		mid = left + (right-left)/2

		if math.Abs(mid-prevMid) < 0.0000001 {
			break
		}
		prevMid = mid

		if check(mid, n, w, h, &words) {
			left = mid
		} else {
			right = mid
		}
	}
	fmt.Fprintf(out, "%.6f\n", left)
}

func check(k float64, n int, w, h float64, words *[]word) bool {
	var currentW, currentH float64
	for i := range n {
		wordWidth := k * (*words)[i].w
		wordHeight := k * (*words)[i].h

		if wordWidth > w || wordHeight > h {
			return false
		}

		if i == 0 {
			currentW = wordWidth
			currentH = wordHeight

			continue
		}

		if (*words)[i-1].h == (*words)[i].h && currentW+wordWidth <= w {
			currentW += wordWidth
		} else if currentH+wordHeight <= h {
			currentW = wordWidth
			currentH += wordHeight
		} else {
			return false
		}

	}
	return true
}
