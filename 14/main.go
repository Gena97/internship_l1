// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{}

	v = 1337
	fmt.Printf("Тип переменной: %s\n", getType(v))

	v = "gg"
	fmt.Printf("Тип переменной: %s\n", getType(v))

	v = true
	fmt.Printf("Тип переменной: %s\n", getType(v))

	v = make(chan int)
	fmt.Printf("Тип переменной: %s\n", getType(v))
}

func getType(v interface{}) string {
	return reflect.TypeOf(v).String()
}
