package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		var input string
		fmt.Fscan(in, &input)
		fmt.Println(process(input))
		fmt.Println("-")
	}
}

const (
	left  = 'L'
	right = 'R'
	up    = 'U'
	down  = 'D'
	home  = 'B'
	end   = 'E'
	enter = 'N'
)

type cursor struct {
	line int
	pos  int
}
type console struct {
	cursor *cursor
	text   []string
}

func newConsole() *console {
	return &console{&cursor{0, 0}, []string{""}}
}

func (c *console) lineLen() int {
	return len(c.text[c.cursor.line])
}

func (c *console) correctPosition() {
	if c.cursor.pos >= c.lineLen() {
		c.cursor.pos = c.lineLen()
	}
}

func (c *console) enterSymbol(b byte) {
	valid := b >= '0' && b <= '9' || b >= 'a' && b <= 'z'
	if !valid {
		return
	}
	line := c.text[c.cursor.line]
	c.text[c.cursor.line] = line[:c.cursor.pos] + string(b) + line[c.cursor.pos:]
	c.cursor.pos++
}

func (c *console) process(b byte) {
	switch b {
	case left:
		if c.cursor.pos > 0 {
			c.cursor.pos--
		}
	case right:
		if c.cursor.pos < c.lineLen() {
			c.cursor.pos++
		}
	case up:
		if c.cursor.line > 0 {
			c.cursor.line--
			c.correctPosition()
		}
	case down:
		if c.cursor.line < len(c.text)-1 {
			c.cursor.line++
			c.correctPosition()
		}
	case home:
		c.cursor.pos = 0
	case end:
		c.cursor.pos = c.lineLen()
	case enter:
		firstHalf := c.text[c.cursor.line][:c.cursor.pos]
		secondHalf := c.text[c.cursor.line][c.cursor.pos:]
		c.text = append(c.text[:c.cursor.line+1], c.text[c.cursor.line:]...)
		c.text[c.cursor.line] = firstHalf
		c.cursor.line++
		c.text[c.cursor.line] = secondHalf
		c.cursor.pos = 0
	default:
		c.enterSymbol(b)
	}
}

func process(input string) string {
	c := newConsole()
	for i := 0; i < len(input); i++ {
		c.process(input[i])
	}
	return strings.Join(c.text, "\n")
}
