package main

// Задача 3. Генератор чисел
// Напиши функцию gen(n int) <-chan int, которая отдаёт числа 0..n-1 и закрывает канал. main читает через for range.

// Тренирует: паттерн generator, defer close.

func Gen(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}