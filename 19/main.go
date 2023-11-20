// 19. Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	"strings"
)

func reverseString(inputStr string) string {
	// Разбиваем строку на массив символов
	charArray := strings.Split(inputStr, "")

	// Переворачиваем массив символов
	for i := 0; i < len(charArray)/2; i++ {
		charArray[i], charArray[len(charArray)-1-i] = charArray[len(charArray)-1-i], charArray[i]
	}

	// Объединяем символы обратно в строку
	reversedStr := strings.Join(charArray, "")

	return reversedStr
}

func main() {
	inputStr := "главрыба — абырвалг"
	reversedStr := reverseString(inputStr)
	fmt.Println(reversedStr)
}
