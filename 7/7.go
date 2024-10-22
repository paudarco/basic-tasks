package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.

// syncMap - структура для конкурентной записи данных в map
type syncMap struct {
	mu   sync.RWMutex
	data map[string]string
}

// конструктор для syncMap
func newSyncMap() *syncMap {
	return &syncMap{data: make(map[string]string)}
}

// добавляет новую пару ключ-значение в map
func (sm *syncMap) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// получает значение по ключу из map
func (sm *syncMap) Get(key string) (string, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

// удаляет пару ключ-значение из map
func (sm *syncMap) Remove(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func main() {
	sm := newSyncMap()
	var wg sync.WaitGroup
	wg.Add(2)

	// запускаем 2 горутины
	// одна заполняет мапу значениями от 1 до 10,
	// другая - от 11 до 20
	go func() {
		for i := 1; i <= 10; i++ {
			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("value-%d", i)
			sm.Set(key, value)
			v, ok := sm.Get(key)
			if !ok {
				fmt.Println("Key not found")
			} else {
				fmt.Printf("Value for key 'key-%d': %s\n", i, v)
			}
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()
	// для имитации задержки вызываем time.Sleep каждую итерацию
	go func() {
		for i := 11; i <= 20; i++ {
			key := fmt.Sprintf("key-%d", i-10)
			value := fmt.Sprintf("value-%d", i)
			sm.Set(key, value)
			v, ok := sm.Get(key)
			if !ok {
				fmt.Println("Key not found")
			} else {
				fmt.Printf("Value for key 'key-%d': %s\n", i-10, v)
			}
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
	// выводим все значения из мапы
	for key, value := range sm.data {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// удаляем пару ключ-значение с ключом "key-7"
	sm.Remove("key-7")
	// проверяем удаление
	v, ok := sm.Get("key-7")
	if ok {
		fmt.Printf("Value for key 'key-7': %s\n", v)
	} else {
		fmt.Println("Key 'key-7' not found")
	}
}
