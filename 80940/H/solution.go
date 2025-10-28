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

	win := make([]bool, n+1)
	win[0] = false
	for i := 1; i <= n; i++ {
		var flag bool
		for try := 1; try <= 3 && try <= i; try++ {
			if i-try == 0 || !isPrime(i-try) {
				if i-try == 0 || !win[i-try] {
					flag = true
					break
				}
			}
		}
		win[i] = flag
	}

	if win[n] {
		fmt.Fprintln(out, 1)

		return
	}

	fmt.Fprintln(out, 2)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
