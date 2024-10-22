package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»).
// Символы могут быть unicode.

func reverse(s string) string {
	// Преобразуем строку в массив rune, чтобы изменять символы вне зависимости от их кодировки.
	strArr := []rune(s)
	// Меняем символы местами с использованием двух указателей.
	// первый - на место последнего, а последний - на место первого
	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return string(strArr)
}

func main() {
	input := "главрыба"
	reversed := reverse(input)
	fmt.Println(reversed) // "главрыба — абырвалг"

	input = "Привет, мир!"
	reversed = reverse(input)
	fmt.Println(reversed) // "!рим, тевирП"
}
