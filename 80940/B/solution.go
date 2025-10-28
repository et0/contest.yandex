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

	var s string
	fmt.Fscan(in, &s)
	// fmt.Println(s)

	var switcher byte = 'L'
	counter := 0
	for left, right := 0, 0; left < len(s); left = right {

		countR, countL, countB := 0, 0, 0
		for ; right < len(s) && (s[right] == 'B' || s[left] == s[right]); right++ {
			switch s[right] {
			case 'L':
				countL++
			case 'R':
				countR++
			case 'B':
				countB++
			}
		}

		// fmt.Printf("switcher=%s left=%d right=%d countR=%d countL=%d countB=%d\n", string(switcher), left, right, countR, countL, countB)

		if right == len(s) {
			switch switcher {
			case 'L':
				if countL+countB+1 > 1+countR+countB {
					counter += 1 + countR
				} else {
					counter += countL + 1
				}
			case 'R':
				if 1+countL+countB+1 > countR+countB {
					counter += countR
				} else {
					counter += 1 + countL + 1
				}
			}
		} else {
			switch switcher {
			case 'L':
				if 1+countR+countB >= countL+countB {
					counter += countL
				} else {
					counter += 1 + countR
					switcher = 'R'
				}
			case 'R':
				if 1+countL+countB >= countR+countB {
					counter += countR
				} else {
					counter += 1 + countL
					switcher = 'L'
				}
			}
		}
		counter += countB

		// fmt.Println(string(switcher), counter, "\n")
	}

	fmt.Fprintln(out, counter)
}
