package main

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
// func main() {
//   someFunc()
// }

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func createHugeString(size int) string {
	// b - буфер для эффективного сложения символов
	var b strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "界")
	}

	return b.String()
}

func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	// руна может занимать не один байт
	fmt.Println(utf8.RuneLen('界')) // 3

	// в данном случае мы срезаем по количеству байт, а не по количеству рун
	// так как символ, например иероглиф в нашем случае, может занимать 3 байт,
	// в итоге получим срез из 33 символов (в 100 байтах 33 символа по 3 байта)
	justString = v[:100]
	fmt.Println(utf8.RuneCountInString(justString))

	// преобразовываем строку в слайс рун
	r := []rune(v)
	// в даннам случае мы срезаем по количеству рун
	justString = string(r[:100])
	fmt.Println(utf8.RuneCountInString(justString))
}
func main() {
	someFunc()
}
