package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// TODO: распарсить два числа, посчитать tip и total, вывести с 2 знаками после запятой
	_ = fmt.Sprintf
	_ = os.Args
	_ = strconv.ParseFloat

	if len(os.Args) != 3 {
		fmt.Println("Неверное количество аргументов")
		os.Exit(1)
	}

	bill, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Неверный формат входных данных")
		os.Exit(1)
	}

	percent, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Неверный формат входных данных")
		os.Exit(1)
	}

	tip := bill * percent / 100.0
	total := bill + tip

	fmt.Printf("Чаевые: %.2f, всего: %.2f\n", tip, total)
}
