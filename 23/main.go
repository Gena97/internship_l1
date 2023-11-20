// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	s = deleteElementFromSlice(s, 2)

	fmt.Println(s)
}

func deleteElementFromSlice(s []int, n int) []int {
	if n > len(s)-1 || n < 0 {
		fmt.Println("Номер элемента выходит за границы массива")
		return s
	}
	return append(s[:n], s[n+1:]...)
}
