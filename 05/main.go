// 5. Разработать программу, которая будет последовательно отправлять значения в канал,
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
	go receiverData(dataChannel)

	go senderData(dataChannel)

	// Ждем N секунд
	time.Sleep(N * time.Second)

	// Закрываем канал после N секунд
	close(dataChannel)

	fmt.Println("Завершение работы")
}

func senderData(ch chan int) {
	for i := 1; ; i++ {
		ch <- i
		fmt.Printf("Отправлено значение: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func receiverData(ch chan int) {
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Получено значение:", data)
		}
	}
}
