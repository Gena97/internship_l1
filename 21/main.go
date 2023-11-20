// 21. Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Исходный интерфейс, несовместимый с целевым интерфейсом
type LegacyPrinter interface {
	Print(string) string
}

// Целевой интерфейс
type Printer interface {
	PrintMessage(string)
}

// Адаптер структура, реализующая целевой интерфейс
type Adapter struct {
	LegacyPrinter
}

// Преобразование вызова интерфейса целевого интерфейса к вызову исходного интерфейса
func (a *Adapter) PrintMessage(message string) {
	fmt.Println(a.Print(message))
}

// Объект, имеющий исходный несовместимый интерфейс
type MyLegacyPrinter struct{}

// Реализация метода исходного интерфейса
func (p *MyLegacyPrinter) Print(s string) string {
	return fmt.Sprintf("Legacy Printer: %s", s)
}

func main() {
	legacyPrinter := &MyLegacyPrinter{}
	adapter := &Adapter{legacyPrinter}

	adapter.PrintMessage("Hello, World!")
}
