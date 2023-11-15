// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

// Задаем значение N в секундах
const N = 5

func main() {
	// Создаем канал для передачи данных
	dataChannel := make(chan int)

	// Запускаем горутину для отправки значений в канал
	go sendData(dataChannel)

	// Ждем N секунд
	time.Sleep(N * time.Second)

	// Закрываем канал после N секунд
	close(dataChannel)

	fmt.Println("Главная горутина завершает работу")
}

func sendData(ch chan int) {
	for i := 1; ; i++ {
		// Отправляем значение в канал
		select {
		case ch <- i:
			fmt.Printf("Отправлено значение: %d\n", i)
		case <-time.After(1 * time.Second):
			// Выход из цикла, если канал закрыт
			fmt.Println("Канал закрыт. Остановка sendData.")
			return
		}
	}
}
