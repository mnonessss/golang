package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// TODO: проверить количество аргументов
	// TODO: распарсить a, op, b
	// TODO: выполнить операцию и вывести результат
	_ = fmt.Sprintf
	_ = os.Args
	_ = strconv.Atoi

	if len(os.Args) != 4 {
		fmt.Println("Неверное количество аргументов")
		os.Exit(2)
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Первое значение должно быть числом")
		os.Exit(1)
	}

	op := os.Args[2]

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Первое значение должно быть числом")
		os.Exit(1)
	}

	var res int

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль")
			os.Exit(1)
		}
		res = a / b
	default:
		os.Exit(1)
	}

	fmt.Printf("Результат: %d\n", res)

}
