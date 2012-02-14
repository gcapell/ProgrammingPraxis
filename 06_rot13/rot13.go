package main

import (
	"log"
	"strings"
	"unicode"
)

const (
	SRC = "Cebtenzzvat Cenkvf vf sha!"
	ROT = 13
)

func main() {
	log.SetFlags(0) // quieter logging
	log.Printf("%s -> %s", SRC, strings.Map(rot13, SRC))

}

func rot13(c rune) rune {
	switch {
	case unicode.IsLetter(c + ROT):
		return c + ROT
	case unicode.IsLetter(c - ROT):
		return c - ROT
	}
	return c
}
