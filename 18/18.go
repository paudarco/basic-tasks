package main

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// счетчик с использованием мьютекса
type Counter struct {
	mu      *sync.RWMutex
	counter int
}

// конструктор
func newCounter() *Counter {
	return &Counter{
		mu:      &sync.RWMutex{},
		counter: 0,
	}
}

// метод увеличения счетчика
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// метод получения текущего значения счетчика
func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter
}

// =================================================================

// счетчик с использованием атомарных операций
type atomicCounter struct {
	counter int64
}

// конструктор
func newAtomicStruct() *atomicCounter {
	return &atomicCounter{
		counter: 0,
	}
}

// метод увеличения счетчика
func (a *atomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

// метод получения текущего значения счетчика
func (a *atomicCounter) Get() int64 {
	return atomic.LoadInt64(&a.counter)
}

func main() {
	wg := &sync.WaitGroup{}
	workersNum := 15
	iterations := 1000
	counter := newCounter()

	// запуск воркеров, увеличивающих счетчик
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				counter.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("counter with mutex scored %d\nnumber of workers: %d\neach worker made %d iterations\n",
		counter.Get(), workersNum, iterations)

	fmt.Println("======================================")

	atomicCounter := newAtomicStruct()
	atomicIterations := 1500

	// запуск воркеров, увеличивающих счетчик
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < atomicIterations; j++ {
				atomicCounter.Inc()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("atomic counter scored %d\nnumber of workers: %d\neach worker made %d iterations\n",
		atomicCounter.Get(), workersNum, atomicIterations)
}
