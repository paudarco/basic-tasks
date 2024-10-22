package main

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	strArr := strings.ToLower(s)
	m := make(map[rune]struct{})
	for _, r := range strArr {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println("Is unique \"Hello\": ", isUnique("Hello"))
	fmt.Println("Is unique \"World\": ", isUnique("World"))
	fmt.Println("Is unique \"World\": ", isUnique("aBcDeAbCdE"))
}
