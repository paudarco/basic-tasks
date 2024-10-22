package main

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

import (
	"fmt"
)

// do принимает интерфейсный параметр i и выводит его тип с помощью приведения типа через switch
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v has type int\n", v)
	case string:
		fmt.Printf("%q has type string\n", v)
	case bool:
		fmt.Printf("%v has type bool\n", v)
	case chan int:
		fmt.Printf("%v has type chan int\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var v interface{}

	// вывод типа данных через форматирование
	v = 10
	fmt.Printf("Type of v: %T\n", v)
	do(v)

	v = "hello"
	fmt.Printf("Type of v: %T\n", v)
	do(v)

	v = true
	fmt.Printf("Type of v: %T\n", v)
	do(v)

	v = make(chan int)
	fmt.Printf("Type of v: %T\n", v)
	do(v)

}
