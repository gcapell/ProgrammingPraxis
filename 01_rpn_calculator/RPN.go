package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.ReadFrom(os.Stdin)

	x := bytes.Fields(b.Bytes())

	for i, j := range x {
		fmt.Printf("a: %v, b:%v\n", i, string(j))
	}

}
