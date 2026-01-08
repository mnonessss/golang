package main

import "fmt"

// RotateMatrix поворачивает матрицу на 90 градусов по часовой стрелке.
func RotateMatrix(matrix [][]int) {
	// Реализуй меня

	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		// {1, 2, 3},
		// {4, 5, 6},
		// {7, 8, 9},
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println("Исходная матрица:")
	printMatrix(matrix)

	RotateMatrix(matrix)

	fmt.Println("После поворота:")
	printMatrix(matrix)
	// Ожидаемый вывод:
	// [7 4 1]
	// [8 5 2]
	// [9 6 3]
}
