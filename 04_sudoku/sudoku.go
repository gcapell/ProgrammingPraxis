package main

import (
	"log"
)

const (
	puzzle1 =   `
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


func main() {
	log.SetFlags(0) // quieter logging

	log.Print(puzzle1)
}
