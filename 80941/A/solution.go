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

	var a, b, S int
	fmt.Fscan(in, &a, &b, &S)

	D := (a-b)*(a-b) + 4*S
	sqrt := int(math.Sqrt(float64(D)))
	if sqrt*sqrt != D {
		fmt.Fprintln(out, -1)
		return
	}

	L1 := (a + b + sqrt) / 2
	L2 := (a + b - sqrt) / 2
	if L1 > 0 && L1 > b && L1 > a && (a+b+sqrt)%2 == 0 {
		fmt.Fprintln(out, L1)
	} else if L2 > 0 && L2 > b && L2 > a && (a+b-sqrt)%2 == 0 {
		fmt.Fprintln(out, L2)
	} else {
		fmt.Fprintln(out, -1)
	}
}
