package main

import (
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
	var b bytes.Buffer
	b.ReadFrom(os.Stdin)
	log.SetFlags(0) // quieter logging

	stack := make([]float64, 0, 50)

	for _, sym := range bytes.Fields(b.Bytes()) {

		op, ok := operators[sym[0]]
		if ok {
			tos := len(stack)
			newValue := op(stack[tos-2], stack[tos-1])
			log.Printf("%.2f %s %.2f = %.2f", stack[tos-2], string(sym[0]), stack[tos-1], newValue)
			stack[tos-2] = newValue
			stack = stack[:tos-1]
		} else {
			s := string(sym)
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Panic(s, err)
			}
			fmt.Println(f)
			stack = append(stack, f)
		}
		fmt.Println("stack:", stack)

	}

}
