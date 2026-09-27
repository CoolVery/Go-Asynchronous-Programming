package main

import (
	"time"
)

// Задача 6. Таймаут через select
// Напиши функцию WaitFor(ch <-chan int, d time.Duration) (int, bool), которая ждёт значение из канала не дольше d. Возвращает (value, true) или (0, false) при таймауте.

// Тренирует: select + time.After.

func WaitFor(ch <-chan int, d time.Duration) (int, bool) {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case val := <-ch:
		return val, true
	case <-timer.C:
		return 0, false
	}
}