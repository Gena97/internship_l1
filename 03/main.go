// 3. Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var totalSum int64 // Общая сумма квадратов

func calculateAndSumSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	square := int64(number * number)

	atomic.AddInt64(&totalSum, square) // Добавляем квадрат числа к общей сумме
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		go calculateAndSumSquare(number, &wg)
	}

	wg.Wait()

	// Выводим общую сумму квадратов
	fmt.Printf("Сумма квадратов: %d\n", totalSum)
}
