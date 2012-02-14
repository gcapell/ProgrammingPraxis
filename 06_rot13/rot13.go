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
	log.Printf("%s -> %s", SRC, rot13(SRC))

}

func rot13(s string) string {
	return strings.Map(func(c rune) rune {
		switch {
		case unicode.IsLetter(c + ROT):
			return c + ROT
		case unicode.IsLetter(c - ROT):
			return c - ROT
		}
		return c
	}, s)
}
