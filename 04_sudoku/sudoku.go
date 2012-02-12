package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
)

const (
	SIZE        = 9
	NSQUARES    = SIZE * SIZE
	SQUARE_SIZE = 3

	puzzle1 = `
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

	puzzle2 = `
	003 020 600
	900 305 001
	001 806 400
	
	008 102 900
	700 000 008
	006 708 200
	
	002 609 500
	800 203 009
	005 010 300`
)

type (
	Board [NSQUARES]uint16
	Pos   struct{ row, col int }
)

func main() {
	log.SetFlags(0) // quieter logging

	var b Board

	log.Print("test", (5/3)*3, (6/3)*3)
	b.LoadFrom(puzzle2)
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

func (b *Board) Assign(p Pos, n uint16) bool {
	pos := p.pos()
	if b[pos]&1<<n == 0 {
		return false
	}
	b[pos] = 1 << n

	for _, peer := range peers(p) {
		if !b.Eliminate(peer, n) {
			return false
		}
	}
	return true
}

func peers(p Pos) []Pos {

	reply := make([]Pos, 0, 24)

	for j := 0; j < SIZE; j++ {

		// row
		if j != p.row {
			reply = append(reply, Pos{j, p.col})
		}

		// col
		if j != p.col {
			reply = append(reply, Pos{p.row, j})
		}

		// square?
		top := (p.row / SQUARE_SIZE) * SQUARE_SIZE
		left := (p.col / SQUARE_SIZE) * SQUARE_SIZE
		for r := 0; r < SQUARE_SIZE; r++ {
			for c := 0; c < SQUARE_SIZE; c++ {
				p2 := Pos{top + r, left + c}
				if p2 != p {
					reply = append(reply, p2)
				}
			}
		}

	}
	return reply
}

// Does this bitmask represent a single bit?
// If so, which one?
var SINGLEVALUES = make(map[uint16]uint16)

func init() {
	for j := uint16(1); j <= 9; j++ {
		SINGLEVALUES[1<<j] = j
	}
}

func (b *Board) Eliminate(p Pos, n uint16) bool {
	pos := p.pos()
	if b[pos]&1<<n == 0 {
		return true // already eliminated
	}
	b[pos] &= ^(1 << n)

	// If we're left with a single value, use it
	if n2, ok := SINGLEVALUES[b[pos]]; ok {
		if !b.Assign(p, n2) {
			return false
		}
	}
	return true
	// FIXME - more propagation!
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

	var cellStrings [NSQUARES]string
	width := 0
	for j := 0; j < NSQUARES; j++ {
		cs := cellString(b[j])
		cellStrings[j] = cs
		if len(cs) > width {
			width = len(cs)
		}
	}
	var buf bytes.Buffer

	for j := 0; j < NSQUARES; j++ {
		buf.WriteString(fmt.Sprintf("%*s", width+1, cellStrings[j]))
		if j%3 == 2 {
			buf.WriteString(fmt.Sprintf(" | "))
		}
		if j%9 == 8 {
			buf.WriteString(fmt.Sprintf("\n"))
		}
		if j%27 == 26 {
			buf.WriteString(fmt.Sprintf("\n"))
		}
	}
	return buf.String()
}

func (b *Board) LoadFrom(s string) {
	for j := 0; j < NSQUARES; j++ {
		b[j] = 0x3ff
	}
	var p Pos
	for _, c := range s {
		if c == '.' || c == '0' {
			p.Next()
		}
		if c >= '1' && c <= '9' {
			n, _ := strconv.Atoi(string(c))
			b.Assign(p, uint16(n))
			p.Next()
		}
	}
}
