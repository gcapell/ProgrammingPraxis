package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

type op_fn func(float64, float64) float64

var operators = map[byte]op_fn{
	'+': func(a, b float64) float64 { return a + b },
	'-': func(a, b float64) float64 { return a - b },
	'*': func(a, b float64) float64 { return a * b },
	'/': func(a, b float64) float64 { return a / b },
}

func main() {
	var b bytes.Buffer
	b.ReadFrom(os.Stdin)
	log.SetFlags(0) // quieter logging

	stack := make([]float64, 0, 50)
	stack = process(b.Bytes(), stack)
}

func process(line []byte, stack []float64) []float64 {
	for _, sym := range bytes.Fields(line) {

		op, ok := operators[sym[0]]
		if ok {
			tos := len(stack)
			stack[tos-2] = op(stack[tos-2], stack[tos-1])
			stack = stack[:tos-1]
		} else {
			s := string(sym)
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Panic(s, err)
			}
			stack = append(stack, f)
		}
		log.Print(stack)
	}
	return stack
}
