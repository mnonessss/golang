package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n float64
	if len(os.Args) != 2 {
		fmt.Println("You entered incorrect number of parameters")
		os.Exit(2)
	}
	n, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("You entered not float parameter")
		os.Exit(2)
	}
	fmt.Printf("%.1f\n", n)
}
