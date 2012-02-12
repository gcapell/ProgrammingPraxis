package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
)

const (
	SIZE     = 9
	NSQUARES = SIZE * SIZE
	puzzle1  = `
	4.. ... 8.5
	.3. ... ...
	... 7.. ...
	
	.2. ... .6.
	... .8. 4..
	... .1. ...
	
	... 6.3 .7.
	5.. 2.. ...
	1.4 ... ...
	`
)

type (
	Board [NSQUARES]uint16
	Pos   struct{ row, col int }
)

func main() {
	log.SetFlags(0) // quieter logging

	var b Board

	b.LoadFrom(puzzle1)
	log.Print(&b)
}

func (p *Pos) Next() {
	p.col += 1
	if p.col == SIZE {
		p.col = 0
		p.row += 1
	}
}

func (p *Pos) pos() int {
	return p.row*SIZE + p.col
}

func (b *Board) Assign(p Pos, n uint16) {
	b[p.pos()] = 1 << n
}

func cellString(c uint16) string {
	if c == 0 {
		return "."
	}

	s := ""
	for j := uint8(1); j <= 9; j++ {
		if c&(1<<j) != 0 {
			s += fmt.Sprintf("%d", j)
		}
	}
	return s
}

func (b *Board) String() string {

	var buf bytes.Buffer
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			buf.WriteString(cellString(b[row*SIZE+col]) + " ")
			if col%3 == 2 {
				buf.WriteString(" ")
			}
		}
		buf.WriteString("\n")
		if row%3 == 2 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

func (b *Board) Init() {

}

func (b *Board) LoadFrom(s string) {
	for j := 0; j < NSQUARES; j++ {
		b[j] = 0x3ff
	}
	var p Pos
	for _, c := range s {
		if c == '.' {
			p.Next()
		}
		if c >= '1' && c <= '9' {
			n, _ := strconv.Atoi(string(c))
			b.Assign(p, uint16(n))
			p.Next()
		}
	}
}
