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

	if n <= 2 {
		fmt.Fprintln(out, n)
		return
	}

	one, two, three := 1, 1, 2
	for i := 3; i <= n; i++ {
		next := one + two + three
		one, two, three = two, three, next
	}

	fmt.Fprintln(out, three)

}
