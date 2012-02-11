package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.ReadFrom(os.Stdin)

	for _, sym := range bytes.Fields(b.Bytes()) {

			switch(sym[0]) {
				case '+': fmt.Println("plus")
				case '-': fmt.Println("minus")
				case '*': fmt.Println("multiply")
				case '/': fmt.Println("divide")
				default: fmt.Println(string(sym))
			}

	}

}
