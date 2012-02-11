package main

import (
	"bufio"
	"bytes"
	"fmt"
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
	log.SetFlags(0) // quieter logging

	r := bufio.NewReader(os.Stdin)
	stack := make([]float64, 0, 50)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		stack = process(line, stack)
		log.Print(stack[len(stack)-1])
	}

}

func process(line []byte, stack []float64) []float64 {
	for _, sym := range bytes.Fields(line) {

		op, ok := operators[sym[0]]
		if ok {
			tos := len(stack)
			if tos < 2 {
				fmt.Fprintf(os.Stderr, "Stack %v too small for %s\n",
					stack, string(sym[0]))
				return stack
			}
			stack[tos-2] = op(stack[tos-2], stack[tos-1])
			stack = stack[:tos-1]
		} else {
			s := string(sym)
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return stack
			}
			stack = append(stack, f)
		}
		// log.Print(stack)
	}
	return stack
}
