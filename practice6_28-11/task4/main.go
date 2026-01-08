package main

import (
	"fmt"
	"strings"
)

/*

Создайте систему валидации для разных типов полей формы.

Определите интерфейс Validatable с методом Validate() (bool, string), где:

bool - результат валидации

string - сообщение об ошибке (пустая строка если ошибок нет)

Реализуйте структуры:

EmailField с полем Value string

PasswordField с полями Value string и MinLength int

NumberField с полями Value int и Min, Max int

Для каждой структуры реализуйте свою логику валидации:

Email: проверка на наличие @ и точки

Password: проверка минимальной длины

Number: проверка диапазона

Создайте функцию ValidateForm(fields []Validatable) []string, которая возвращает все сообщения об ошибках.

Продемонстрируйте работу с формой, содержащей все типы полей.

*/

type Validatable interface {
	Validate() (bool, string)
}

type EmailField struct {
	Value string
}

func (field EmailField) Validate() (bool, string) {
	if !strings.Contains(field.Value, "@") || !strings.Contains(field.Value, ".") {
		return false, "email not contain '@' or '.'"
	}

	return true, ""
}

type PasswordField struct {
	Value     string
	MinLength int
}

func (field PasswordField) Validate() (bool, string) {
	if len(field.Value) < field.MinLength {
		return false, "password is too short"
	}
	return true, ""
}

type NumberField struct {
	Value, Min, Max int
}

func (field NumberField) Validate() (bool, string) {
	if field.Value < field.Min || field.Max < field.Value {
		return false, "number is not valid"
	}
	return true, ""
}

func ValidateForm(fields []Validatable) []string {
	var errors []string
	for _, f := range fields {
		res, err := f.Validate()
		if !res {
			errors = append(errors, err)
		}
	}

	return errors
}

func main() {
	form := []Validatable{
		// EmailField{"invalid-email"},
		// PasswordField{Value: "123", MinLength: 6},
		// NumberField{Value: 150, Min: 0, Max: 100},
		EmailField{"valid@email.com"},
		PasswordField{Value: "strongpassword", MinLength: 6},
		NumberField{Value: 50, Min: 0, Max: 100},
	}

	fmt.Println("Валидация формы...")
	errors := ValidateForm(form)

	if len(errors) == 0 {
		fmt.Println("✅ Все поля валидны!")
	} else {
		fmt.Println("❌ Найдены ошибки:")
		for _, err := range errors {
			fmt.Println(" -", err)
		}
	}

	// Демонстрация индивидуальной валидации
	fmt.Println("\nИндивидуальная проверка:")
	email := EmailField{"test@example.com"}
	if isValid, msg := email.Validate(); isValid {
		fmt.Println("✅ Email валиден")
	} else {
		fmt.Println("❌", msg)
	}
}
