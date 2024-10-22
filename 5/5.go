package main

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// writer - горутина, бесконечно записывающая значения в канал кажду секунду
// при отмене контекста завершает работу
func writer(ch chan<- int, ctx context.Context) {
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("writer finished")
			return
		default:
			fmt.Printf("writing: %d\n", i)
			ch <- i
		}
		time.Sleep(time.Second)
	}
}

// reader читает значения из канала ch,
// пока он не закроется
func reader(ch <-chan int) {
	for v := range ch {
		fmt.Printf("reading: %d\n", v)
	}
	fmt.Println("reader finished")
}

func main() {
	var ch = make(chan int)
	var wg sync.WaitGroup
	var seconds int

	// цикл спрашивает количество секунд работы программы
	// и продолжает спрашивать, пока введенное значение
	// не окажется единственным, больше 0 и меньше 60 (условно)
	for {
		fmt.Printf("\nHow many seconds program should work: ")
		n, err := fmt.Scan(&seconds)
		if err != nil || n <= 0 {
			continue
		}
		fmt.Println()
		if n <= 0 {
			fmt.Println("Number of seconds must be positive")
			continue
		}
		if n > 60 {
			fmt.Println("Too long, insert again")
			continue
		}

		break
	}
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// при завершении writer закрывается главный канал
	// и автоматически завершает работу и reader
	go func() {
		writer(ch, ctx)
		close(ch)
		wg.Done()
	}()
	go func() {
		reader(ch)
		wg.Done()
	}()

	time.AfterFunc(time.Duration(seconds)*time.Second, func() {
		cancel()
	})
	wg.Wait()

	fmt.Println("program finished")
}
