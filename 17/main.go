// 17. Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	array := []int{1, 2, 4, 7, 10, 32, 55, 77, 98}

	binnarySearch(array, 8)

}

func binnarySearch(array []int, n int) int {
	lenArray := len(array)

	start, end := 0, lenArray-1

	for start <= end {
		mid := start + (end-start)/2

		if array[mid] == n {
			fmt.Println("Значение найдено", mid)
			return mid
		}

		if array[mid] < n {
			start = mid + 1
		} else if array[mid] > n {
			end = mid - 1
		}
	}

	fmt.Println("Значение не найдено")
	return -1
}
