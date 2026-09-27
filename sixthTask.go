package main

// Задача 6. Неблокирующее чтение
// Напиши функцию TryReceive(ch <-chan int) (int, bool), которая пытается прочитать из канала, но не блокируется, если данных нет.

// Тренирует: select с default.
// Подсказка: select { case v := <-ch: ...; default: ... }.

func TryReceive(ch <-chan int) (int, bool) {
	select {
	//Проверка, что чтение из НЕ закрытого канала
	case val, ok := <-ch:
		if ok {
			return val, true
		} else {
			return 0, false
		}
	//Из-за default канал не будет блокироваться и продолжит свою работу
	default:
		return 0, false
	}
}