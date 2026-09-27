package main

import (
	"fmt"
	"sync"
)

func main() {
	//First task
	//PingPong()

	//SecondTask
	//Функция возвращает канал из которого сразу читаем значение
	//get := <-SumInCh([]int{1, 2, 65, 3, 4, 5})
	//fmt.Println(get)

	//ThirdTask
	// for num := range Gen(3) {
	// 	fmt.Println(num)
	// }

	//FourthTask
	// Запусти 5 горутин, каждая пишет свой ID в общий канал. main собирает все 5 значений и печатает. 
	// Используй sync.WaitGroup для закрытия канала после всех писателей.
	// var wg sync.WaitGroup

	// out := make(chan int)
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// Передаем в горутину i как айди
	// 	go func(num int) {
	// 		defer wg.Done()
	// 		out <- num
	// 	}(i)
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(out)
	// }()
	// for id := range out {
	// 	fmt.Println(id)
	// }
	
}