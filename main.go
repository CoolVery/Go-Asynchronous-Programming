package main

import (
	"fmt"
	_ "sync"
	_ "time"
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
	
	//FiftTask
	// ch1 := make(chan int, 1)
    // ch1 <- 42
    // v, ok := WaitFor(ch1, 1*time.Second)
    // fmt.Println("случай 1:", v, ok) // 42 true

    // // Случай 2: таймаут — никто не пишет
    // ch2 := make(chan int)
    // v, ok = WaitFor(ch2, 100*time.Millisecond)
    // fmt.Println("случай 2:", v, ok) // 0 false

	//SixTask
	// 1. Есть значение
    // ch1 := make(chan int, 1)
    // ch1 <- 42
    // v, ok := TryReceive(ch1)
    // fmt.Println("случай 1 (есть значение):", v, ok) // 42 true

    // // 2. Пусто
    // ch2 := make(chan int)
    // v, ok = TryReceive(ch2)
    // fmt.Println("случай 2 (пусто):", v, ok) // 0 false

    // // 3. Закрыт
    // ch3 := make(chan int)
    // close(ch3)
    // v, ok = TryReceive(ch3)
    // fmt.Println("случай 3 (закрыт):", v, ok) // 0 false
}