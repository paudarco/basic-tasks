package main

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a и b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	// создание двух переменных типа big.Int
	a := big.NewInt((rand.Int63n(11) + 1) * 1 << (rand.Int63n(31) + 20))
	b := big.NewInt((rand.Int63n(11) + 1) * 1 << (rand.Int63n(31) + 20))

	fmt.Printf("1st number: %s\n2nd number: %s\n\n", a.String(), b.String())

	// переменная для хранения умножения, деления, сложения и вычитания
	res := big.NewInt(0)
	fmt.Println("multiplex:", res.Mul(a, b))
	fmt.Println("division:", res.Div(a, b))
	fmt.Println("sum:", res.Add(a, b))
	fmt.Println("subtract:", res.Sub(a, b))
}
