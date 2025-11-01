package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Ride struct {
	start int
	end   int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		n, m  int
		buses int
		s     string
	)

	fmt.Fscan(in, &n)
	forward := make([]Ride, n)
	for i := range n {
		fmt.Fscan(in, &s)
		forward[i] = parseRide(s)
	}
	sort.Slice(forward, func(i, j int) bool {
		return forward[i].start < forward[j].start
	})

	fmt.Fscan(in, &m)
	backward := make([]Ride, m)
	for i := range m {
		fmt.Fscan(in, &s)
		backward[i] = parseRide(s)
	}
	sort.Slice(backward, func(i, j int) bool {
		return backward[i].start < backward[j].start
	})

	forwardAvailable := &IntHeap{}
	backwardAvailable := &IntHeap{}
	heap.Init(forwardAvailable)
	heap.Init(backwardAvailable)

	for i, j := 0, 0; i < len(forward) || j < len(backward); {
		if i < len(forward) && (j >= len(backward) || forward[i].start <= backward[j].start) {
			if backwardAvailable.Len() > 0 && (*backwardAvailable)[0] <= forward[i].start {
				heap.Pop(backwardAvailable)
			} else {
				buses++
			}
			heap.Push(forwardAvailable, forward[i].end)
			i++

			continue
		}

		if forwardAvailable.Len() > 0 && (*forwardAvailable)[0] <= backward[j].start {
			heap.Pop(forwardAvailable)
		} else {
			buses++
		}
		heap.Push(backwardAvailable, backward[j].end)
		j++
	}

	fmt.Fprintln(out, buses)
}

func parseRide(s string) Ride {
	parts := strings.Split(s, "-")
	start := timeconv(parts[0])
	end := timeconv(parts[1])

	return Ride{start, end}
}

func timeconv(timeStr string) int {
	parts := strings.Split(timeStr, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])

	return hours*60 + minutes
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
