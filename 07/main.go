package main

import (
	"errors"
	"fmt"
	"sync"
)

type SafeNumbers struct {
	sync.RWMutex
	numbers map[int]int
}

func (sn *SafeNumbers) Add(num int) {
	sn.Lock()
	defer sn.Unlock()
	if sn.numbers == nil {
		sn.numbers = make(map[int]int)
	}
	sn.numbers[num] = num
}

func (sn *SafeNumbers) Get(num int) (int, error) {
	sn.RLock()
	defer sn.RUnlock()
	if number, ok := sn.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Number does not exist")
}

func addNumbersMap(sn *SafeNumbers, amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < amount; i++ {
		sn.Add(i)
	}
}

func main() {
	safeNumbers := &SafeNumbers{
		numbers: make(map[int]int),
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go addNumbersMap(safeNumbers, 100, &wg)
	go addNumbersMap(safeNumbers, 300, &wg)

	wg.Wait()

	fmt.Println(safeNumbers.numbers)
}
