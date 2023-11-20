// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	stopGoroutine1()
	stopGoroutine2()
	stopGoroutine3()
}

// Остановка горутины путем отправки сигнала в канал.
func stopGoroutine1() {
	quitChan := make(chan bool)
	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				// Вывод сообщения каждые 3 секунды
				fmt.Println("Это тестовая горутина 1")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	// Задержка для вывода некоторых сообщений из горутины
	time.Sleep(time.Second * 8)

	// Остановка горутины
	quitChan <- true
	fmt.Println("Горутина 1 остановлена!")

	// Проверка, остановилась ли горутина
	fmt.Println("Это конец функции 1")

	time.Sleep(time.Second * 4)
}

// Закрытие горутины путем закрытия канала.
func stopGoroutine2() {
	var wg sync.WaitGroup
	wg.Add(1)

	strChan := make(chan string)
	go func() {
		for {
			element, ok := <-strChan
			// Проверка, закрыт ли канал
			if !ok {
				println("Горутина 2 остановлена!")
				wg.Done()
				return
			}
			println(element)
		}
	}()
	strChan <- "это"
	strChan <- "тестовая"
	strChan <- "строка"
	strChan <- "сообщения"
	close(strChan)

	// Ожидание остановки всех горутин
	wg.Wait()
	// Вывод последнего сообщения
	fmt.Println("Это конец функции 2!")
	time.Sleep(time.Second * 4)
}

// Закрытие горутины путем отмены контекста.
func stopGoroutine3() {
	// Канал для бесконечного выполнения горутины
	channel := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // Если выполнен cancel()
				channel <- struct{}{}
				fmt.Println("В горутине получен сигнал закрытия контекста")
				return
			default:
				fmt.Println("Это из горутины 3")
			}

			// Вывод сообщения каждую секунду
			time.Sleep(1 * time.Second)
		}
	}(ctx)

	// Горутина для закрытия вышеуказанного контекста
	go func() {
		time.Sleep(5 * time.Second)
		// Закрытие контекста
		cancel()
	}()

	// Бесконечное выполнение горутины
	<-channel
	fmt.Println("Это конец функции 3")
}
