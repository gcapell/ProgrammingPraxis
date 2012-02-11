package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var b bytes.Buffer
	b.ReadFrom(os.Stdin)

	for _, sym := range bytes.Fields(b.Bytes()) {

		switch sym[0] {
		case '+':
			fmt.Println("plus")
		case '-':
			fmt.Println("minus")
		case '*':
			fmt.Println("multiply")
		case '/':
			fmt.Println("divide")
		default:
			s := string(sym)
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Panic(s, err)
			}
			fmt.Println(f)
		}

	}

}
