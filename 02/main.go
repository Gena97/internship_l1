// 2. Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func calculateSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик ожидающих горутин при завершении функции

	square := number * number
	fmt.Println(square)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // Создаем WaitGroup для ожидания завершения всех горутин

	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик ожидающих горутин
		go calculateSquare(number, &wg)
	}

	wg.Wait() // Ожидаем завершения всех горутин
}
