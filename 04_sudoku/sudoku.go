package main

// Based entirely on http://norvig.com/sudoku.html

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

	// Apparently hard
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

	// Can be solved without backtracking
	simple_puzzle = `
	..3 .2. 6..
	9.. 3.5 ..1
	..1 8.6 4..
	
	..8 1.2 9..
	7.. ... ..8
	..6 7.8 2..
	
	..2 6.9 5..
	8.. 2.3 ..9
	..5 .1. 3..`

	programming_praxis_puzzle = `
	7.. 1.. ...
	.2. ... .15
	... ..6 39.
	
	2.. .18 ...
	.4. .9. .7.
	... 75. ..3
	
	.78 5.. ...
	56. ... .4.
	... ..1 ..2
	`
)

type (
	Board [NSQUARES]uint16
	Pos   struct{ row, col int }
)

func main() {
	log.SetFlags(0) // quieter logging

	var b Board

	b.LoadFrom(programming_praxis_puzzle)
	log.Print(search(&b))
}

func search(b *Board) *Board {
	if b.solved() {
		return b
	}
	p := b.minChoicePos()
	for _, n := range bits(b[p.pos()]) {
		c := *b // copy
		if c.Assign(p, n) {
			attempt := search(&c)
			if attempt != nil {
				return attempt
			}
		}
	}
	return nil
}

func bits(n uint16) []uint16 {
	reply := make([]uint16, 0)
	for j := uint16(1); j <= 9; j++ {
		if n&(1<<j) != 0 {
			reply = append(reply, j)
		}
	}
	return reply
}

func (b *Board) solved() bool {
	for _, n := range b {
		_, ok := SINGLEVALUES[n]
		if !ok {
			return false
		}
	}
	return true
}

// Return position within board which
// has minimal (>1) number of choices
func (b *Board) minChoicePos() Pos {
	minset := 9
	minPos := 0
	for pos, n := range b {
		set := bitsSet(n)
		switch {
		case set == 1:
			continue
		case set == 2:
			return nToPos(pos)
		case set < minset:
			minset = set
			minPos = pos
		}
	}
	return nToPos(minPos)
}

func bitsSet(n uint16) int {
	c := 0
	for ; n != 0; c++ {
		n &= n - 1 // Clear least significant bit
	}
	return c
}

func nToPos(n int) Pos {
	return Pos{n / SIZE, n % SIZE}
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
	if b[pos]&(1<<n) == 0 {
		return false
	}
	b[pos] = 1 << n

	for _, unit := range units(p) {
		for _, peer := range unit {
			if !b.Eliminate(peer, n) {
				return false
			}
		}
	}
	return true
}

func units(p Pos) [][]Pos {
	reply := make([][]Pos, 3)
	for j := 0; j < 3; j++ {
		reply[j] = make([]Pos, 0, 8)
	}

	for j := 0; j < SIZE; j++ {

		// row
		if j != p.row {
			reply[0] = append(reply[0], Pos{j, p.col})
		}

		// col
		if j != p.col {
			reply[1] = append(reply[1], Pos{p.row, j})
		}
	}

	// square?
	top := (p.row / SQUARE_SIZE) * SQUARE_SIZE
	left := (p.col / SQUARE_SIZE) * SQUARE_SIZE
	for r := 0; r < SQUARE_SIZE; r++ {
		for c := 0; c < SQUARE_SIZE; c++ {
			p2 := Pos{top + r, left + c}
			if p2 != p {
				reply[2] = append(reply[2], p2)
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
	if b[pos]&(1<<n) == 0 {
		return true // already eliminated
	}
	b[pos] &= ^(1 << n)

	// If we're left with a single value, use it
	if n2, ok := SINGLEVALUES[b[pos]]; ok {
		if !b.Assign(p, n2) {
			return false
		}
	}

	// For each unit of p, if there's one remaining
	// place to put n, do that.
	for _, u := range units(p) {
		switch nFound, firstPos := findInUnit(b, u, n); nFound {
		case 0:
			// no location in this unit. contradiction
			return false
		case 1:
			// Exactly one location. use it.
			if !b.Assign(firstPos, n) {
				return false
			}
		}
	}

	return true
}

func findInUnit(b *Board, u []Pos, n uint16) (int, Pos) {
	found := false
	var foundPos Pos
	for _, p := range u {
		if b[p.pos()]&(1<<n) != 0 {
			if found {
				// >1 locations in this unit.
				return 2, foundPos
			} else {
				found = true
				foundPos = p
			}
		}
	}
	if !found {
		return 0, foundPos
	}
	return 1, foundPos
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
		b[j] = 0x3fe
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
