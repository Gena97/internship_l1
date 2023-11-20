// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package main

import "fmt"

func main() {
	createSetStrings1()
	createSetStrings2()
}

func createSetStrings1() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	stringSet := make(map[string]bool)

	for _, str := range sequence {
		stringSet[str] = true
	}

	fmt.Println(stringSet)
}

// В данном примере используется минимум памяти
func createSetStrings2() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	stringSet := make(map[string]struct{})

	for _, str := range sequence {
		stringSet[str] = struct{}{}
	}

	fmt.Println(stringSet)
}
