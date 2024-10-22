package main

// Удалить i-ый элемент из слайса.

import "fmt"

func remove(slice []int, i int) []int {
	// создаем результирующий слайс из объединения значений слева и справа от искомого
	result := append(slice[:i], slice[i+1:]...)
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Before: ", numbers)
	numbers = remove(numbers, 2)
	fmt.Println("After: ", numbers) // Output: [1 2 4 5]
}
