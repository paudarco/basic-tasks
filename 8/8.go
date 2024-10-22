package main

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

import "fmt"

func main() {
	var num int64
	var bit, index int

	// Ввод числа, индекса и значения бита
	for {
		fmt.Print("Number, bit index, bit value (x y z): ")
		fmt.Scan(&num, &index, &bit)

		if bit > 1 || bit < 0 || index > 64 {
			fmt.Println("Invalid input. Bit value must be 0 or 1, and bit index must be less than 64.")
			continue
		}
		break
	}

	// Создаем число-маску
	// Сдвигаем 1 на позицию index и выполняем побитовое ИЛИ с num, если bit равен 1
	// В противном случае выполняем побитовое ИСКЛЮЧАЮЩЕЕ ИЛИ с num, сдвинутым на позицию index
	if bit == 1 {
		num |= (1 << uint(index))
	} else {
		num &^= (1 << uint(index))
	}

	fmt.Printf("Result: %d\n", num)
}
