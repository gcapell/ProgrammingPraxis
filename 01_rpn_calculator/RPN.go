package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.ReadFrom(os.Stdin)
	fmt.Printf("Hello %v!\n", b.Len())

}
