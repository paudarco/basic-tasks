package main

// Реализовать пересечение двух неупорядоченных множеств.

import "fmt"

func main() {
	set1 := []int{12, 8, 4, 5, 2, 10, 7}
	set2 := []int{3, 15, 9, 6, 8, 2, 5}

	// intersection - мапа с числом в ключе и количеством повторений в значении
	intersection := make(map[int]int)
	// result - массив совпадающих чисел из обоих множеств
	result := make([]int, 0)

	for _, v := range set1 {
		intersection[v] += 1
	}

	for _, v := range set2 {
		intersection[v] += 1
	}

	// если у числа больше 1 повторений - пересечение
	for num, count := range intersection {
		if count > 1 {
			result = append(result, num)
		}
	}

	fmt.Println("Intersection:", result)
}
