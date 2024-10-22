package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import "fmt"

// быстрая сортировка Хоара
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// pivot - случайная отправная точка для сравнения
	// less и greater - больше и меньше pivot
	pivot := arr[0]
	var less, greater []int

	for _, val := range arr[1:] {
		if val <= pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	// в результат вкладываем отсортированное less (меньше pivot/середины)
	// и greater (больше pivot/середины) вместе с pivot по центру
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func main() {
	fmt.Println(quickSort([]int{9, 50, 32, 65, 8, 10, 31, 98, 42, 5, 7}))
}
