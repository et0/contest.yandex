package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Node struct {
	value  string
	left   *Node
	right  *Node
	height int
	width  int
	canvas [][]byte
}

type App struct {
	symbols []string
	pos     int
}

func NewApp(s string) *App {
	symbols := processing(s)

	return &App{symbols: symbols}
}

func processing(s string) []string {
	var symbols []string
	var current []rune

	for _, ch := range s {
		if unicode.IsSpace(ch) {
			continue
		}

		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^' || ch == '(' || ch == ')' {
			if len(current) > 0 {
				symbols = append(symbols, string(current))
				current = nil
			}
			symbols = append(symbols, string(ch))
		} else {
			current = append(current, ch)
		}
	}

	if len(current) > 0 {
		symbols = append(symbols, string(current))
	}

	return symbols
}

func (p *App) getTreeFromExpression() *Node {
	node := p.parseTerm()

	for p.pos < len(p.symbols) {
		token := p.symbols[p.pos]
		if token == "+" || token == "-" {
			p.pos++
			right := p.parseTerm()
			node = &Node{value: token, left: node, right: right}
		} else if token == ")" {
			break
		} else {
			p.pos++
		}
	}

	return node
}

func (p *App) parseTerm() *Node {
	node := p.parseFactor()

	for p.pos < len(p.symbols) {
		token := p.symbols[p.pos]
		if token == "*" || token == "/" {
			p.pos++
			right := p.parseFactor()
			node = &Node{value: token, left: node, right: right}
		} else if token == "+" || token == "-" || token == ")" {
			break
		} else {
			p.pos++
		}
	}

	return node
}

func (p *App) parseFactor() *Node {
	node := p.parseElement()

	if p.pos < len(p.symbols) && p.symbols[p.pos] == "^" {
		p.pos++
		right := p.parseFactor()
		node = &Node{value: "^", left: node, right: right}
	}

	return node
}

func (p *App) parseElement() *Node {
	if p.pos >= len(p.symbols) {
		return nil
	}

	token := p.symbols[p.pos]
	p.pos++

	if token == "(" {
		node := p.getTreeFromExpression()
		if p.pos < len(p.symbols) && p.symbols[p.pos] == ")" {
			p.pos++
		}
		return node
	}

	return &Node{value: token}
}

func (n *Node) prepare() {
	if n.left == nil && n.right == nil {
		n.height = 1
		n.width = len(n.value)
		return
	}

	if n.left != nil {
		n.left.prepare()
	}
	if n.right != nil {
		n.right.prepare()
	}

	leftH, leftW := 0, 0
	rightH, rightW := 0, 0

	if n.left != nil {
		leftH = n.left.height
		leftW = n.left.width
	}
	if n.right != nil {
		rightH = n.right.height
		rightW = n.right.width
	}

	n.height = max(leftH, rightH) + 2
	n.width = leftW + rightW + 5
}

func (n *Node) create() {
	if n.canvas != nil {
		return
	}

	n.prepare()

	n.canvas = make([][]byte, n.height)
	for i := range n.canvas {
		n.canvas[i] = make([]byte, n.width)
		for j := range n.canvas[i] {
			n.canvas[i][j] = ' '
		}
	}

	if n.left == nil && n.right == nil {
		for i, ch := range n.value {
			n.canvas[0][i] = byte(ch)
		}
		return
	}

	leftH, leftW := 0, 0
	rightH, rightW := 0, 0

	if n.left != nil {
		n.left.create()
		leftH = n.left.height
		leftW = n.left.width

		// копируем левый
		for i := 0; i < leftH; i++ {
			for j := 0; j < leftW; j++ {
				n.canvas[i+2][j] = n.left.canvas[i][j]
			}
		}

		// рисуем левый
		startLeftPrint := 0
		if n.left.left != nil {
			startLeftPrint = n.left.left.width + 2
		}
		n.canvas[0][startLeftPrint] = '.'
		n.canvas[1][startLeftPrint] = '|'

		for j := startLeftPrint + 1; j <= leftW; j++ {
			n.canvas[0][j] = '-'
		}
	}
	if n.right != nil {
		n.right.create()
		rightH = n.right.height
		rightW = n.right.width

		// копируем правый
		rightStart := leftW + 5
		for i := 0; i < rightH; i++ {
			for j := 0; j < rightW; j++ {
				n.canvas[i+2][rightStart+j] = n.right.canvas[i][j]
			}
		}

		// рисуем правые
		rightLineStart := leftW + 4
		rightLineEnd := n.width - 1
		if n.right.right != nil {
			rightLineEnd -= n.right.right.width + 2
		}
		for j := rightLineStart; j < rightLineEnd; j++ {
			n.canvas[0][j] = '-'
		}
		n.canvas[0][rightLineEnd] = '.'
		n.canvas[1][rightLineEnd] = '|'
	}

	n.canvas[0][leftW+1] = '['
	n.canvas[0][leftW+2] = byte(n.value[0])
	n.canvas[0][leftW+3] = ']'
}

func (n *Node) print(out *bufio.Writer) {
	n.create()
	for _, row := range n.canvas {
		fmt.Fprintln(out, string(row))
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	if s == "" {
		return
	}

	parser := NewApp(s)
	head := parser.getTreeFromExpression()
	head.print(out)
}
