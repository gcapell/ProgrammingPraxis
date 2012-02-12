package main

import (
	"log"
	"math/rand"
)

const (
	BOARDSIZE          = 5
	MIDDLE             = 2
	NUMBERS            = 75
	NUMBERS_PER_COLUMN = NUMBERS / BOARDSIZE
)

type Board struct {
	// counts of how many squares are filled in
	// each row/col/diagonal
	rows      [BOARDSIZE]int
	cols      [BOARDSIZE]int
	diagonals [2]int
}

type PlaceList []*int

type Game struct {
	boards []*Board
	places [NUMBERS]PlaceList
}

func (g *Game) Init(boards int) {
	for j := 0; j < boards; j++ {
		b := new(Board)
		b.Init(g)
		g.boards = append(g.boards, b)
	}
}

func main() {
	log.SetFlags(0) // quieter logging

	average(500, 100000)
}

func average(boards, iterations int) {
	var g Game

	g.Init(boards)
	log.Printf("%d boards, %d iterations, %.2f average",
		boards, iterations, g.findAverage(iterations))

}

func (g *Game) addPointer(n int, p *int) {
	g.places[n] = append(g.places[n], p)
}

func (g *Game) findAverage(iterations int) float64 {
	var length = 0
	for j := 0; j < iterations; j++ {
		length += g.gameLength()
		g.reset()
	}
	return float64(length) / float64(iterations)
}

func (g *Game) reset() {
	for _, b := range g.boards {
		b.reset()
	}
}

func (b *Board) reset() {
	b.diagonals[0] = 1
	b.diagonals[1] = 1
	for j := 0; j < BOARDSIZE; j++ {
		var val = 0
		if j == MIDDLE {
			val = 1
		}
		b.cols[j] = val
		b.rows[j] = val
	}
}

func (g *Game) gameLength() int {
	for pos, n := range rand.Perm(NUMBERS) {
		for _, p := range g.places[n] {
			*p += 1
			if *p == BOARDSIZE {
				// log.Print("gameLenth:", pos)
				return pos
			}
		}
	}
	panic("cannot get here")
	return 0
}

func (b *Board) Init(g *Game) {
	b.reset()

	for col := 0; col < BOARDSIZE; col++ {
		perms := rand.Perm(NUMBERS_PER_COLUMN)[:BOARDSIZE]
		for row, n := range perms {
			if row == MIDDLE && col == MIDDLE {
				continue
			}
			n += col * NUMBERS_PER_COLUMN
			g.addPointer(n, &b.cols[col])
			g.addPointer(n, &b.rows[row])
			if row == col {
				g.addPointer(n, &b.diagonals[0])
			}
			if row+col == BOARDSIZE-1 {
				g.addPointer(n, &b.diagonals[1])
			}
		}
	}
}
