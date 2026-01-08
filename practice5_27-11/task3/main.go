package main

import "fmt"

// MergeMaps объединяет две карты, выбирая максимальное значение для дублирующихся ключей.
func MergeMaps(map1, map2 map[string]int) map[string]int {
	// Реализуй меня

	result := make(map[string]int)

	for key, value := range map1 {
		result[key] = value
	}

	for key, _ := range map2 {
		if map1[key] >= map2[key] {
			result[key] = map1[key]
		} else {
			result[key] = map2[key]
		}
	}

	return result
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "e": 2}
	m2 := map[string]int{"b": 5, "c": 1, "d": 4}

	merged := MergeMaps(m1, m2)
	fmt.Println("Объединенная карта:", merged)
	// Ожидаемый вывод: map[a:1 b:5 c:3 d:4]
}
