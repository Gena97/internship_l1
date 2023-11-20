/* 15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

// Ответ: если "v" < 100, то произойдет выход за граинцы массива, поэтому нужна проверка.

package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		fmt.Println("Длина значения v меньше 100")
	}
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	var builder strings.Builder
	builder.Grow(size) // Выделение памяти под строку заранее

	for i := 0; i < size; i++ {
		builder.WriteString("x") // Добавление символов к строке
	}

	return builder.String() // Получение результирующей строки
}
