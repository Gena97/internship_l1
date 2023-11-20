// 18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	// Запускаем несколько горутин, которые инкрементируют счетчик
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println(counter.GetCount())
}
