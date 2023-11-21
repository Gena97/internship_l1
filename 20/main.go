// 20. Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Разбиваем строку на слова
	words := strings.Split(s, " ")

	// Меняем слова местами
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Соединяем и возвращаем новую строку
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun - sun dog snow"
	reversed := reverseWords(s)
	fmt.Println(reversed)
}
