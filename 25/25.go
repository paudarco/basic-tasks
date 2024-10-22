package main

// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func sleep(i int) {
	// time.after возвращает канал, который будет блокировать продолжение работы
	// функции (в нашем случае будет блокировка ее завершения) до тех пор,
	// пока из него не будет прочитано отправленное значение
	<-time.After(time.Second * time.Duration(i))
}

func main() {
	now := time.Now()
	sleep(5) // При вызове sleep(5) функция будет ждать 5 секунд
	fmt.Println("Execution time:", time.Since(now))
}
