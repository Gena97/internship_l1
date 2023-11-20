// 25. Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

// Функция mySleep реализует задержку выполнения программы на указанное количество секунд
func mySleep(seconds int) {
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		return
	}
}

func main() {
	fmt.Println("Задержка на 5 секунд")
	mySleep(5)
	fmt.Println("Прошло 5 секунд")
}
