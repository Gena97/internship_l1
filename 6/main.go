// 6. Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Пример 1: использование канала для остановки горутины
	quitChan := make(chan struct{})
	go testGoroutineQuit(quitChan)

	time.Sleep(time.Second * 10)
	close(quitChan) // закрываем канал
	fmt.Println("Горутина 1 завершает работу")
	time.Sleep(time.Second * 5)

	// Пример 2: использование sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go testGoroutineWaitGroup(&wg)
	time.Sleep(time.Second * 2)
	wg.Done() // сигнализируем о завершении горутины
	fmt.Println("Горутина 2 завершает работу")
	time.Sleep(time.Second * 5)

	// Пример 3: использование контекста для управления горутиной
	ctx, cancel := context.WithCancel(context.Background())
	go testGoroutineContext(ctx)
	time.Sleep(time.Second * 5)
	cancel() // отменяем контекст
	time.Sleep(time.Second * 2)
	fmt.Println("Горутина 3 завершает работу")
}

func testGoroutineQuit(quitChan <-chan struct{}) {
	for {
		select {
		case <-quitChan:
			fmt.Println("Получен сигнал завершения горутины 1")
			return
		default:
			fmt.Println("Это testGoroutineQuit")
			time.Sleep(time.Second * 3)
		}
	}
}

func testGoroutineWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println("Это testGoroutineWaitGroup")
		time.Sleep(time.Second * 3)
	}
}

func testGoroutineContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Произошла отмена контекста горутины 3")
			return
		default:
			fmt.Println("Это testGoroutineContext")
			time.Sleep(time.Second * 3)
		}
	}
}
