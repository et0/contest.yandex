package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tree struct {
	tree   []int
	length int
}

func NewTree(length int) *Tree {
	return &Tree{make([]int, length+1), length}
}

func (t *Tree) update(index int, delta int) {
	for i := index; i <= t.length; i += i & -i {
		t.tree[i] += delta
	}
}

func (t *Tree) query(index int) int {
	sum := 0
	for i := index; i > 0; i -= i & -i {
		sum += t.tree[i]
	}
	return sum
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	fmt.Fscan(in, &n, &x)

	maxSize := n + 100000
	queue := make([]int, maxSize)
	tree := NewTree(maxSize)

	for i := range n {
		fmt.Fscan(in, &queue[i])
		if queue[i] >= x {
			tree.update(i+1, 1)
		}
	}

	var m, event, a, k int
	fmt.Fscan(in, &m)
	start, end := 0, n
	for range m {
		fmt.Fscan(in, &event)

		if event == 1 {
			fmt.Fscan(in, &a)
			queue[end] = a
			if a >= x {
				tree.update(end+1, 1)
			}
			end++

			continue
		}

		if event == 2 {
			if queue[start] >= x {
				tree.update(start+1, -1)
			}
			start++

			continue
		}

		if event == 3 {
			fmt.Fscan(in, &k)

			count := tree.query(start+k) - tree.query(start)
			fmt.Fprintln(out, count)
		}
	}
}
