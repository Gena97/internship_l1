// 26. Разработать программу, которая проверяет,
// что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

package main

import "fmt"

func main() {

	checker("abcd")
	checker("abCdefAaf")
	checker("aabcd")

}

func checker(s string) bool {
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
