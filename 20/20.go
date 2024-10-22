package main

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	// Преобразовать строку в массив слов, разделенных пробелом
	strArr := strings.Split(s, " ")
	// меняем местами первый и последний элементы по порядку
	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return strings.Join(strArr, " ")
}

func main() {
	s := "Hello World!"
	reversed := reverse(s)
	fmt.Println(reversed) // Output: World! Hello

	s = "snow dog sun"
	reversed = reverse(s)
	fmt.Println(reversed) // Output: sun dog snow
}
