// 20. Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "главрыба — абырвалг"
	reversed := reverseString1(str)
	fmt.Println(reversed)

	reversed = reverseString2(str)
	fmt.Println(reversed)
}

func reverseString1(str string) string {
	var reversedBuilder strings.Builder
	// Разбиваем строку на руны (unicode символы)
	runes := []rune(str)
	// Проходим по рунам в обратном порядке и добавляем их в Builder
	for i := len(runes) - 1; i >= 0; i-- {
		reversedBuilder.WriteRune(runes[i])
	}
	return reversedBuilder.String()
}

func reverseString2(str string) string {
	runes := []rune(str)
	reversed := make([]rune, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		reversed[len(runes)-1-i] = runes[i]
	}
	return string(reversed)
}
