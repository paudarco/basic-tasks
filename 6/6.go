package main

// Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// горутина, завершающая свою работу с отменой контекста
// либо по истечению времени, либо из-за ручной отмены
func doneWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done goroutine closed with context")
			return
		default:
			fmt.Println("goroutine with context is running")
			time.Sleep(1 * time.Second)
		}
	}
}

// горутина, завершающая работу при проверке
// переменой состояния канала
func doneWithClosingChannel(ch <-chan struct{}) {
	for {
		_, ok := <-ch
		if !ok {
			fmt.Println("done with closed channel")
			return
		}
		fmt.Println("goroutine with closing channel is working")
	}
}

// горутина, завершающая работу при закрытии канала
// и невозможности чтения него
func doneWithClosingChannelRange(ch <-chan struct{}) {
	for _ = range ch {
		fmt.Println("goroutine with range of channel working")
	}
	fmt.Println("done with closing range of channel")
}

// горутина, заверающая работу при получении значения
// из отдельного канала для остановки работы
func doneWithStopSignal(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("done goroutine with stop signal")
			return
		default:
			fmt.Println("goroutine with stop signal is running")
			time.Sleep(1 * time.Second)
		}
	}
}

// горутина, завершающая работу при получении значения из канала таймера
func doneWithTimer(timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			fmt.Println("done goroutine with timer")
			return
		default:
			fmt.Println("goroutine with timer is running")
			time.Sleep(1 * time.Second)
		}
	}
}

// горутина, заверщающая работу при получении значения из канала таймера,
// созданного внутри горутины
func doneWithTimeout() {
	timer := time.After(time.Second * 4)
	for {
		select {
		case <-timer:
			fmt.Println("done goroutine with timeout")
			return
		default:
			fmt.Println("goroutine with timeout is running")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// запуск горутин, завершающейся при закрытии канала
	chOk := make(chan struct{})
	wg.Add(2)
	go func() {
		defer wg.Done()
		doneWithClosingChannel(chOk)
	}()
	go func() {
		defer wg.Done()
		doneWithClosingChannelRange(chOk)
	}()
	wg.Add(1)
	time.AfterFunc(time.Second*3, func() {
		wg.Done()
		close(chOk)
	})
	// горутина, отправляющей данные в канал chOk
	// для демонстрации работы
	time.Sleep(time.Second * 1)
	chOk <- *new(struct{})

	//================================================================

	// горутина, завершающаяся при отмене контекста таймаутом
	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		doneWithContext(ctx)
	}()
	//================================================================

	// горутина, завершающаяся при отмене контекста вручную
	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(time.Second*2, cancel)
		doneWithContext(ctx)
	}()

	//================================================================

	// горутина, завершающаяся при отправке любого значения в канал
	// закрытие канала тоже вызовет событие записи в канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch := make(chan struct{})
		time.AfterFunc(time.Second*3, func() {
			close(ch)
		})
		doneWithStopSignal(ch)
	}()
	//================================================================

	// горутина, завершающаяся при истечении таймаута внутри самой горутины
	wg.Add(1)
	go func() {
		defer wg.Done()
		doneWithTimeout()
	}()
	//================================================================

	// горутина, завершающаяся при истечении таймаута, определенного в аргументах
	wg.Add(1)
	go func() {
		defer wg.Done()
		doneWithTimer(time.NewTimer(time.Second * 3))
	}()

	wg.Wait()
	fmt.Println("all goroutines ended their work")
}
