package main

import (
	"bufio"
	"fmt"
	"os"
)

const winRepeat int = 5

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	if n < 5 && m < 5 {
		fmt.Fprint(out, "No")
		return
	}

	var line string
	data := make([][]byte, n)
	for i := range n {
		fmt.Fscan(in, &line)
		data[i] = []byte(line)
	}

	directs := [5][2]int{
		{-1, 0},  // left
		{-1, -1}, // left-top
		{0, -1},  // top
		{1, -1},  // right-top
		{1, 0},   // right
	}

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if data[y][x] == '.' {
				continue
			}
			for _, d := range directs {
				if !check(x, y, n, m, &data, &d) {
					continue
				}

				fmt.Fprint(out, "Yes")
				return
			}
		}
	}

	fmt.Fprint(out, "No")
}

func check(x, y, n, m int, data *[][]byte, direct *[2]int) bool {
	//fmt.Printf("\n%s %d %d\n", string((*data)[y][x]), x, y)
	for i := 1; i < winRepeat; i++ {
		if x+direct[0]*i < 0 || x+direct[0]*i > m-1 || y+direct[1]*i < 0 || y+direct[1]*i > n-1 {
			//fmt.Println("false1", x+direct[0]*i, y+direct[1]*i)
			return false
		}

		if (*data)[y][x] != (*data)[y+direct[1]*i][x+direct[0]*i] {
			//fmt.Println("false2", string((*data)[y+direct[1]*i][x+direct[0]*i]))
			return false
		}

		//fmt.Println("true", string((*data)[y+direct[1]*i][x+direct[0]*i]))
	}

	//fmt.Println(x, y)

	return true
}
