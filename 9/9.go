package main

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

import "fmt"

// arrGen принимает массив чисел, возвращает входной канал только для чтения
// и в отдельной горутине отправляет числа из массива в канал
func arrGen(nums ...int) <-chan int {
	in := make(chan int)
	go func() {
		for _, num := range nums {
			in <- num
		}
		close(in)
	}()
	return in
}

// sqrOut принимает входной канал и возвращает выходной канал для чтения
// параллельно в горутине отправляя квадраты чисел из входного канала в канал выхода
func sqrOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}
func main() {
	in := arrGen(2, 4, 5, 1, 12, 8, 9)
	out := sqrOut(in)

	// получаем квадраты чисел из канала выхода
	for num := range out {
		fmt.Printf("%d ", num)
	}
}
