// 26. Разработать программу, которая проверяет,
// что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func main() {

	checker1("abcd")
	checker1("abCdefAaf")
	checker1("aabcd")

	checker2("abcd")
	checker2("abCdefAaf")
	checker2("aabcd")

}

func checker1(s string) bool {
	var charSet [26]bool

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v -= 32
		}
		if v >= 'A' && v <= 'Z' {
			if charSet[v-'A'] {
				fmt.Println("Строка неуникальна")
				return false
			}
			charSet[v-'A'] = true
		}
	}

	fmt.Println("Строка уникальна")
	return true
}

func checker2(s string) bool {
	s = strings.ToLower(s) // Приведение всех символов строки к нижнему регистру

	charSet := make(map[rune]bool)

	for _, v := range s {
		if charSet[v] {
			fmt.Println("Строка неуникальна")
			return false
		}
		charSet[v] = true
	}

	fmt.Println("Строка уникальна")
	return true
}
