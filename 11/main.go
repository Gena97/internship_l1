// 11. Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func main() {
	set1 := []int{1, 2, 2, 3, 4, 4, 5, 5}
	set2 := []int{2, 3, 4, 4, 5, 6, 7, 8, 5}

	intersection := make([]int, 0)
	set1Count := make(map[int]int)

	for _, el := range set1 {
		set1Count[el]++
	}

	for _, el := range set2 {
		if count, ok := set1Count[el]; ok && count > 0 {
			intersection = append(intersection, el)
			set1Count[el]--
		}
	}

	fmt.Println(intersection)
}
