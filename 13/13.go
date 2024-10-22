package main

// Поменять местами два числа без создания временной переменной.

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Printf("a = %d; b = %d\n", a, b)

	a, b = b, a // Меняем значения переменных a и b

	fmt.Printf("a = %d; b = %d\n", a, b)
}
