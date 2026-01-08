package main

import "fmt"

// FilterPositive возвращает новый срез только с положительными числами.
func FilterPositive(numbers []int) []int {
	// Реализуй меня
	var result []int
	for _, val := range numbers {
		if val > 0 {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	input := []int{-5, 10, 0, 3, -2, 8}
	result := FilterPositive(input)
	fmt.Println("Исходный срез:", input)
	fmt.Println("Только положительные:", result)
	// Ожидаемый вывод: [10 3 8]
}
