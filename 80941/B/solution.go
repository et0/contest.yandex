package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	ways        []int
	firstVisit  bool
	returnVisit bool
	dist        int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var a, b int
	nodes := make([]Node, n)
	for range n - 1 {
		fmt.Fscan(in, &a, &b)

		nodes[a-1].ways = append(nodes[a-1].ways, b-1)
		nodes[b-1].ways = append(nodes[b-1].ways, a-1)
	}

	queue := make([]int, 0, n)
	for i, n := range nodes {
		if len(n.ways) > 1 {
			continue
		}

		queue = append(queue, i)
		nodes[i].firstVisit = true
	}

	minDist := 100000
	for len(queue) > 0 {
		current := queue[0]

		nodes[current].returnVisit = true

		for _, i := range nodes[current].ways {
			// если следующий узел ещё ниразу не посещался, то добавляем его в очередь на посещение
			if !nodes[i].firstVisit {
				nodes[i].firstVisit = true
				nodes[i].dist = nodes[current].dist + 1

				queue = append(queue, i)

				continue
			}

			// если следующий узел будет посещен повторно (сразу же вернулись обратно), то пропускаем его
			if nodes[i].returnVisit {
				continue
			}
			nodes[i].returnVisit = true

			// дистанция меньше минимальной
			if nodes[current].dist+nodes[i].dist+1 < minDist {
				minDist = nodes[current].dist + nodes[i].dist + 1
			}

		}

		queue = queue[1:]
	}

	fmt.Fprintln(out, minDist)
}
