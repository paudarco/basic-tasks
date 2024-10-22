package main

// Реализовать паттерн «адаптер» на любом примере.

import "fmt"

// группа структур, которые мы должны адаптировать
// просто собака - adaptee адаптируемый класс
type Dog struct{}

// реакция собаки
func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав.")
}

// просто кошка - adaptee адаптируемый класс
type Cat struct{}

// реакция кошки
func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Мяу-мяу.")
	}
}

// целевой интерфейс, который должен работать в системе - Target в документации
type AnimalAdapter interface {
	Operation()
}

// адаптер собаки
type DogAdapter struct {
	*Dog
}

// newDogAdapter - конструктор адаптера собаки
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// operation - адаптированный метод под наш целевой интерфейс
func (adapter *DogAdapter) Operation() {
	adapter.WoofWoof()
}

// адаптер кошки
type CatAdapter struct {
	*Cat
}

// operation - адаптированный метод под наш целевой интерфейс
func (adapter *CatAdapter) Operation() {
	adapter.MeowMeow(true)
}

// newCatAdapter - конструктор адаптера кошки
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {
	fmt.Println("Подходишь к двери, а там")
	myFamily := []AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{})}
	for _, member := range myFamily {
		member.Operation()
	}
}
