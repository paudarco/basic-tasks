package main

// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println("temperatires: ", temperatures)
	groupedTemperatures := make(map[int][]float64)

	for _, t := range temperatures {
		// делим темпераатуру на 10, чтобы перенести первый разряд числа
		// в дробную часть и избавиться от него при приведении типа
		group := int(t/10) * 10
		groupedTemperatures[group] = append(groupedTemperatures[group], t)
	}

	for group, temperatures := range groupedTemperatures {
		fmt.Printf("%d: %v\n", group, temperatures)
	}
}
