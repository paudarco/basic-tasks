package main

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

// squareToChan вычисляет квадрат числа и отправляет его в канал c
func squareToChan(x int, c chan<- int) {
	c <- x * x
}

// squareToChanResult вычисляет квадраты чисел в массиве arr
// в этом примере квадрат числа отправляется в результирующий канал
func squareToChanResult() {
	var arr = [5]int{2, 4, 6, 8, 10}
	// создание канала c
	var c chan int = make(chan int)
	var wg sync.WaitGroup

	for _, x := range arr {
		// Вызов функции squareToChan в отдельной горутине и добавление её в WaitGroup
		wg.Add(1)
		go func(a int) {
			squareToChan(x, c)
			wg.Done()
		}(x)
	}

	go func() {
		wg.Wait()
		close(c) // Закрытие канала, чтобы не ожидать больше результатов
	}()

	for x := range c {
		fmt.Printf("%d ", x)
	}

	fmt.Println("ToChanResult done")
}

// squareFromChanResult вычисляет квадраты числа из канала in
// и отправляет результат в канал out
func squareFromChan(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func squareFromChanResult() {
	var arr = [5]int{2, 4, 6, 8, 10}
	var in chan int = make(chan int)
	var out chan int = make(chan int)
	var wg sync.WaitGroup

	go squareFromChan(in, out)

	for _, i := range arr {
		wg.Add(1)
		go func(num int) {
			in <- num
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(in) // Закрытие канала, чтобы не ожидать больше результатов
	}()

	for i := range out {
		fmt.Printf("%d ", i)
	}

	fmt.Println("FromChanResult done")
}

func main() {
	squareToChanResult()
	squareFromChanResult()
}
