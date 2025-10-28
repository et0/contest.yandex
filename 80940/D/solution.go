package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const maxLength int = 20

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var word string
	fmt.Fscan(in, &word)

	var n int
	fmt.Fscan(in, &n)

	var s string
	dict := make(map[string]int, n)
	for range n {
		fmt.Fscan(in, &s)
		dict[s] = len(s)
	}

	result := make([]string, 0, 100)
	checkNext(0, &result, &word, &dict)

	slices.Reverse(result)

	fmt.Fprintln(out, strings.Join(result, " "))
}

func checkNext(start int, result *[]string, word *string, dict *map[string]int) bool {
	for right := 0; start+right < len(*word) && right < maxLength; right++ {
		length, ok := (*dict)[(*word)[start:start+right+1]]
		if !ok {
			continue
		}

		if start+length == len(*word) || checkNext(start+length, result, word, dict) {
			*result = append(*result, (*word)[start:start+right+1])

			return true
		}
	}

	return false
}
