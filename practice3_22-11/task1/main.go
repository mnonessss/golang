package main

import (
	"fmt"
)

func main() {
	n := 7
	pi := 3.14
	quoted := "go"

	fmt.Printf("n=%d type=%T hex=%x\n", n, n, n)
	fmt.Printf("pi=%.2f quoted=%q\n", pi, quoted)
}
