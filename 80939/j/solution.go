package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	lists := make(map[string][]uint32, 100000)

	reNewList, _ := regexp.Compile(`^List\s([a-z]{1,10})\s=\snew\sList\(([\d\,]+)\)\n$`)
	reSubList, _ := regexp.Compile(`^List\s([a-z]{1,10})\s=\s([a-z]{1,10})\.subList\((\d+)\,(\d+)\)\n$`)
	reSet, _ := regexp.Compile(`^([a-z]{1,10})\.set\((\d+),(\d+)\)\n$`)
	reAdd, _ := regexp.Compile(`^([a-z]{1,10})\.add\((\d+)\)\n$`)
	reGet, _ := regexp.Compile(`^([a-z]{1,10})\.get\((\d+)\)\n$`)

	var n int
	fmt.Fscan(in, &n)

	for n > 0 {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\n" {
			continue
		}

		n--

		// List a = new List(x,y, ...,z)
		match := reNewList.FindStringSubmatch(line)
		if len(match) == 3 {
			key := lists[match[1]]
			if key == nil {
				lists[match[1]] = make([]uint32, 0, n)
			}

			for _, v := range strings.Split(match[2], ",") {
				number, _ := strconv.Atoi(v)
				lists[match[1]] = append(lists[match[1]], uint32(number))
			}

			continue
		}

		// List b = a.subList(from,to)
		match = reSubList.FindStringSubmatch(line)
		if len(match) == 5 {
			from, _ := strconv.Atoi(match[3])
			to, _ := strconv.Atoi(match[4])
			lists[match[1]] = lists[match[2]][from-1 : to]

			continue
		}

		// a.set(i,x)
		match = reSet.FindStringSubmatch(line)
		if len(match) == 4 {
			index, _ := strconv.Atoi(match[2])
			number, _ := strconv.Atoi(match[3])
			lists[match[1]][index-1] = uint32(number)

			continue
		}

		// a.add(x)
		match = reAdd.FindStringSubmatch(line)
		if len(match) == 3 {
			number, _ := strconv.Atoi(match[2])
			lists[match[1]] = append(lists[match[1]], uint32(number))

			continue
		}

		// a.get(i)
		match = reGet.FindStringSubmatch(line)
		if len(match) == 3 {
			index, _ := strconv.Atoi(match[2])
			fmt.Fprintln(out, lists[match[1]][index-1])

			continue
		}

	}
}
