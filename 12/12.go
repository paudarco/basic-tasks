package main

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("set: %v\n", arr)

	// check - для записи существования элемента
	// result - для хранения результата (убранных дублей)
	check := make(map[string]bool)
	result := make([]string, 0)

	// в цикле проходимся по массиву, если для текущего элемента есть такой же в check, пропускаем его,
	// если нет, добавляем в result и устанавливаем в check, что этот элемент уникален
	// в конце получаем массив без дублей и выводим его на экран
	// таким образом мы создаем собственное множество из исходной последовательности
	for _, v := range arr {
		if _, ok := check[v]; ok {
			continue
		}
		result = append(result, v)
		check[v] = true
	}

	fmt.Printf("subset: %v\n", result)
}
