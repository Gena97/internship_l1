// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func setBit(value int64, pos uint, bit int64) int64 {
	mask := int64(1) << pos
	value &= ^mask
	value |= bit << pos
	return value
}

func main() {
	var num int64
	var pos uint
	var bit int64

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	fmt.Print("Введите позицию бита (от 0): ")
	fmt.Scan(&pos)
	if pos >= 64 || pos < 0 {
		fmt.Println("Позиция бита выходит за границы, введити число от нуля до 64, включая эти числа")
		fmt.Scan(&pos)
	}

	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&bit)
	if bit != 0 && bit != 1 {
		fmt.Println("Значение бита должно быть либо 0, либо 1, попробуйте еще раз")
		fmt.Scan(&bit)
	}

	result := setBit(num, pos, bit)
	fmt.Printf("Результат: %d\n", result)
}
