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

	joy := 0
	min, max := 1000, 1
	data := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &data[i])

		if i%2 == 0 {
			if data[i] < min {
				min = data[i]
			}

			joy += data[i]
		} else {
			if data[i] > max {
				max = data[i]
			}

			joy -= data[i]
		}
	}

	// Если максимальный элемент у Маши больше, чем минимальный элемент у Васи, то надо меняться
	if min < max {
		joy -= min * 2
		joy += max * 2
	}

	fmt.Fprintln(out, joy)
}
