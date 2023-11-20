// 20. Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s) // Разделение строки на слова
	reversed := make([]string, len(words))

	for i := 0; i < len(words); i++ {
		reversed[i] = reverseString(words[i]) // Переворачиваем каждое слово
	}

	return strings.Join(reversed, " ") // Объединяем слова обратно в строку
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Переворачиваем строку
	}
	return string(runes)
}

func main() {
	input := "snow dog sun - sun dog snow"
	output := reverseWords(input)
	fmt.Println(output)
}
