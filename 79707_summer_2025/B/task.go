package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Island struct {
	i    int
	j    int
	time int
}

type Queue []Island

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].time < q[j].time
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x any) {
	*q = append(*q, x.(Island))
}

func (q *Queue) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func New() *Queue {
	h := &Queue{}
	heap.Init(h)

	return h
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	// Куча
	q := New()
	// Исходная карта
	sourceMap := make([][]int, n)
	// Результат карта
	resultMap := make([][]int, n)

	for i := range n {
		sourceMap[i] = make([]int, m)
		resultMap[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(in, &sourceMap[i][j])

			if sourceMap[i][j] == 0 {
				resultMap[i][j] = 0 // клетка вода

				heap.Push(q, Island{i, j, 0})
			} else {
				resultMap[i][j] = -1 // клетка является сушей
			}
		}
	}

	routes := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for q.Len() > 0 {
		current := heap.Pop(q).(Island)

		for _, r := range routes {
			ni, nj := current.i+r[0], current.j+r[1]

			if ni >= 0 && ni < n && nj >= 0 && nj < m {
				// Если клетка ещё не обработана
				if resultMap[ni][nj] == -1 {
					resultMap[ni][nj] = max(current.time, sourceMap[ni][nj])
					heap.Push(q, Island{ni, nj, resultMap[ni][nj]})
				}
			}
		}
	}

	for i := range n {
		for j := range m {
			fmt.Fprint(out, resultMap[i][j], " ")
		}
		fmt.Fprint(out, "\n")
	}
}
