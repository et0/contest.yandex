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

	var lengthS, countPart, lengthSub int
	fmt.Fscan(in, &lengthS, &countPart)
	lengthSub = lengthS / countPart

	var s string
	fmt.Fscan(in, &s)

	var subS string
	words := make(map[string][]int, countPart)
	for i := range countPart {
		fmt.Fscan(in, &subS)
		words[subS] = append(words[subS], i+1)
	}

	for i := 0; i < lengthS; i += lengthSub {

		slice := words[s[i:i+lengthSub]]
		count := len(slice)
		fmt.Fprintf(out, "%d ", slice[count-1])
		words[s[i:i+lengthSub]] = slice[0 : count-1]
	}
}
