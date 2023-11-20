// Разработать конвейер чисел. Даны два канала:
// в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numChannel := make(chan int)
	sqrChannel := make(chan int)

	// Запускаем горутину, которая будет считывать числа из массива и передавать их в канал numbers
	go func() {
		array := []int{1, 2, 3, 4, 5}

		for _, num := range array {
			numChannel <- num
		}

		close(numChannel)
	}()

	// Запускаем горутину, которая будет удваивать числа из канала numbers и передавать их в канал doubledNumbers
	go func() {
		for num := range array {
			sqrChannel <- num * num
		}

		close(sqrChannel)
	}()

	// Выводим удвоенные значения чисел в stdout
	for sqrNum := range sqrChannel {
		fmt.Println(sqrNum)
	}
}
