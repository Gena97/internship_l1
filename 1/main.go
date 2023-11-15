// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human // Встраивание структуры Human
}

// Метод Walk реализует действие "ходить"
func (h *Action) Walk() {
	fmt.Printf("%s идет.\n", h.Name)
}

// Метод Speak реализует действие "говорить"
func (h *Action) Speak(message string) {
	fmt.Printf("%s говорит: %s\n", h.Name, message)
}

// Метод Sleep реализует действие "спать"
func (h *Action) Sleep() {
	fmt.Printf("%s спит.\n", h.Name)
}

func main() {
	// Создаем экземпляр структуры Action
	person := Action{
		Human: Human{
			Name: "Иван",
			Age:  30,
		},
	}

	// Вызываем методы
	person.Walk()
	person.Speak("Привет!")
	person.Sleep()
}
