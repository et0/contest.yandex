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

	uniq := make(map[int]int)
	sequence := make([]int, 0, 9)

	for i := 0; i < k; i++ {
		lastDigit := n % 10

		if lastDigit == 0 {
			break
		}

		if k-1-i < 10 {
			n += lastDigit
			continue
		}

		if index, ok := uniq[lastDigit]; ok {
			sum := 0
			for _, v := range sequence[index:] {
				sum += v
			}

			countRepeat := (k - 1 - i) / len(sequence[index:])
			n += countRepeat * sum
			i += len(sequence[index:]) * countRepeat
		} else {
			uniq[lastDigit] = len(sequence)
			sequence = append(sequence, lastDigit)
		}
		n += lastDigit
	}

	fmt.Fprintln(out, n)
}
