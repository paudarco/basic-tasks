package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	var result int64
	var wg sync.WaitGroup
	nums := []int64{2, 4, 6, 8, 10}

	for _, v := range nums {
		wg.Add(1)
		// каждая горутина выполняет потокобезовасное прибавление квадрата числа
		// через атомарные операции
		go func(num int64) {
			defer wg.Done()
			atomic.AddInt64(&result, num*num)
		}(v)
	}

	wg.Wait()

	fmt.Println(result)
}
