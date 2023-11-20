// 22. Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2 в 20  степени

// Ответ: т.к. не указали, на сколько больше, чем 2^20, воспользуемся библиотекой 'math/big'"

package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigInt1 := big.NewInt(42)
	bigInt2 := big.NewInt(42)

	multiplication := multiplyNumbers(bigInt1, bigInt2)
	fmt.Println("Умножение:", multiplication)

	division := divideNumbers(bigInt1, bigInt2)
	fmt.Println("Деление:", division)

	addition := addNumbers(bigInt1, bigInt2)
	fmt.Println("Сложение:", addition)

	subtraction := subtractNumbers(bigInt1, bigInt2)
	fmt.Println("Вычитание:", subtraction)
}

func multiplyNumbers(a, b *big.Int) *big.Int {
	result := big.NewInt(0).Mul(a, b)
	return result
}

func divideNumbers(a, b *big.Int) *big.Float {
	result := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(a), big.NewFloat(0).SetInt(b))
	return result
}

func addNumbers(a, b *big.Int) *big.Int {
	result := big.NewInt(0).Add(a, b)
	return result
}

func subtractNumbers(a, b *big.Int) *big.Int {
	result := big.NewInt(0).Sub(a, b)
	return result
}
