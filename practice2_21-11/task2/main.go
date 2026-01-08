package main

import (
	"fmt"
	"os"
	"strconv"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на 0\n")
	}
	return a / b, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Неверное количество аргументов")
		os.Exit(2)
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		os.Exit(1)
	}

	res, err := div(a, b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
		os.Exit(1)
	}
	fmt.Printf("Результат: %d\n", res)

}
