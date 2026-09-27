package main

import (
	"fmt"
	"sync"
)

// Задача 1. Ping-pong
// Запусти две горутины. Первая отправляет в канал число, вторая получает и печатает. main ждёт завершения.

// Тренирует: базовую отправку/получение, небуферизованный канал.
// Подсказка: make(chan int) + одна горутина-отправитель и main-получатель.

func pingPong() {
	//Создаем waitGroup
	var wg sync.WaitGroup
	//Создаем небуф. канал
	ch := make(chan int)
	//Добавляем в группу 2 горунтины
	wg.Add(2)
	//Отправитель
	go func() {
		//После выполнения удаляем из группы горунтину и закрываем отпрвавителя
		defer wg.Done()
		defer close(ch)
		 ch <- 7
	}()
	//Читатель
	go func() {
		defer wg.Done()
		get := <-ch
		fmt.Println(get)
	}()
	wg.Wait()
}
