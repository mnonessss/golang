package main

// Напишите функцию WordFrequency, которая принимает строку и возвращает карту,
// где ключи — это уникальные слова в строке, а значения — количество их вхождений.
// Игнорируйте регистр и знаки препинания (для простоты, разделителем считается пробел).

import (
	"fmt"
	"strings"
)

func WordFrequency(a string) map[string]int {

	result := make(map[string]int)

	words := strings.Split(a, " ")

	for _, word := range words {
		word = strings.ToLower(word)
		result[word]++
	}

	return result

}

func main() {
	a := "Hello World world"
	result := WordFrequency(a)

	for key, value := range result {
		fmt.Printf("%s: %d\n", key, value)
	}
}
