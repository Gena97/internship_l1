// 4. Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var numWorkers int

	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	dataChannel := make(chan string)
	done := make(chan struct{})

	var wg sync.WaitGroup

	// Инициализация воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg, done)
	}

	// Обработка сигнала Ctrl+C для завершения программы
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalCh
		close(done) // Закрываем канал done при получении сигнала Ctrl+C
	}()

	// Генерация произвольных данных и отправка их в канал
	for i := 1; ; i++ {
		select {
		case <-done:
			break
		default:
			data := fmt.Sprintf("Data %d", i)
			dataChannel <- data
		}
	}

	// Закрытие основных каналов и ожидание завершения воркеров
	close(dataChannel)
	wg.Wait()
	fmt.Println("Главный поток завершает работу.")
}

func worker(id int, dataChannel <-chan string, wg *sync.WaitGroup, done <-chan struct{}) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				// Канал закрыт, завершаем работу воркера
				fmt.Printf("Воркер %d завершает работу.\n", id)
				return
			}
			fmt.Printf("Воркер %d получил: %s\n", id, data)
		case <-done:
			// Получен сигнал завершения, завершаем работу воркера
			fmt.Printf("Воркер %d завершает работу по сигналу.\n", id)
			return
		}
	}
}
