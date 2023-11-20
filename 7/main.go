// Реализовать конкурентную запись данных в map.

package main

import (
	"errors"
	"log"
	"sync"
)

type SafeNumbers struct {
	sync.RWMutex
	numbers map[int]int
}

func (sn SafeNumbers) Add(num int) {
	sn.Lock()
	defer sn.Unlock()
	if sn.numbers == nil {
		sn.numbers = make(map[int]int)
	}
	sn.numbers[num] = num
}

func (sn SafeNumbers) Get(num int) (int, error) {
	sn.RLock()
	defer sn.RUnlock()
	if number, ok := sn.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Number does not exist")
}

func generateNumbersMap(wg *sync.WaitGroup, sn *SafeNumbers) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		sn.Add(i)
	}
}

func readNumbersMap(wg *sync.WaitGroup, sn *SafeNumbers) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		number, err := sn.Get(i)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(number)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	safeNumbers := &SafeNumbers{
		numbers: make(map[int]int),
	}

	wg.Add(1)
	go generateNumbersMap(&wg, safeNumbers)

	wg.Wait()

	wg.Add(1)
	go readNumbersMap(&wg, safeNumbers)

	wg.Wait()
}
