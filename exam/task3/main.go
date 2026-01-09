package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	runes := []rune(word)
	if len(runes) <= 1 {
		return word
	}

	first := runes[0]
	rest := runes[1:]
	for i, j := 0, len(rest)-1; i < j; i, j = i+1, j-1 {
		rest[i], rest[j] = rest[j], rest[i]
	}

	result := make([]rune, len(runes))
	result[0] = first
	copy(result[1:], rest)

	return string(result)
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)

	encryptedWords := make([]string, len(words))
	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	testPhrases := []string{
		"Pepe Schnele is a legend",
		"Hello world",
		"a",
		"Go",
		"Привет мир",
		"Alpha   Beta Gamma",
		"",
	}

	for _, phrase := range testPhrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("Исходная: %-30s → Зашифровано: %s\n", fmt.Sprintf("«%s»", phrase), fmt.Sprintf("«%s»", encrypted))
	}
}
