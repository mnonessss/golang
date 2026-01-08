package main

import (
	"fmt"
	"sort"
)

/*

Вам нужно отсортировать слайс структур Book по разным полям.

Создайте структуру Book с полями Title (string), Author (string) и Year (int).

Реализуйте интерфейс sort.Interface для слайса []Book:

Сортировка по названию (алфавитная)

Сортировка по году издания (от старых к новым)

Напишите функцию PrintBooks(books []Book), которая красиво выводит информацию о книгах.

В main создайте слайс книг, отсортируйте его разными способами и выведите результаты.

*/

type Book struct {
	Title, Author string
	Year          int
}

type ByTitle []Book

func (b ByTitle) Len() int {
	return len(b)
}

func (b ByTitle) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByTitle) Less(i, j int) bool {
	return b[i].Title < b[j].Title
}

type ByYear []Book

func (b ByYear) Len() int {
	return len(b)
}

func (b ByYear) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByYear) Less(i, j int) bool {
	return b[i].Year < b[j].Year
}

func PrintBooks(books []Book) {
	for i, book := range books {
		fmt.Printf("ID%d, Name: %s, Author: %s, Year: %d\n", i+1, book.Title, book.Author, book.Year)
	}
}

func main() {
	books := []Book{
		{"The Go Programming Language", "Alan Donovan", 2015},
		{"Clean Code", "Robert Martin", 2008},
		{"Design Patterns", "Erich Gamma", 1994},
		{"The C Programming Language", "Dennis Ritchie", 1978},
	}

	fmt.Println("Исходный список:")
	PrintBooks(books)

	fmt.Println("Сортировка по названию:")
	sort.Sort(ByTitle(books))
	PrintBooks(books)

	fmt.Println("Сортировка по году:")
	sort.Sort(ByYear(books))
	PrintBooks(books)
}
