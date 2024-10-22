package main

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// worder под номером workerNum читает число из канала и выводит его,
// пока не будет сигнала об отмене общего для всех контекста
func worker(workerNum int, in <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case num := <-in:
			fmt.Printf("Worker %d received: %d\n", workerNum, num)
		}
	}
}

func main() {
	// канал для записи
	var ch = make(chan int)
	var wg sync.WaitGroup
	var numOfWorkers int
	// цикл для ввода количества worker'ов
	// будет работать, пока не введут
	// неотрицательное число или менее 20
	for {
		fmt.Printf("\nHow many workers: ")
		n, err := fmt.Scan(&numOfWorkers)
		if err != nil || n <= 0 {
			continue
		}
		fmt.Println()
		if numOfWorkers <= 0 {
			fmt.Println("Number of workers must be positive")
			continue
		}
		if numOfWorkers > 20 {
			fmt.Println("Please be merciful and insert again")
			continue
		}
		break
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// запускаем worker'ов в отдельных горутинах
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go func(num int) {
			worker(num, ch, ctx)
			wg.Done()
		}(i)
	}

	// горутина с бесконечным циклом записывает числа в главный поток
	// и делает паузу каждую секунду,
	// пока не отменится общий контекст
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
			time.Sleep(time.Second)
		}
	}()

	// по нажатию ctrl+c сигнал регистрируется в канале stop
	// контекст отменяется и ожидается завершение всех горутин для выхода из программы
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down...")
	cancel()
	wg.Wait()
	close(ch)
}
