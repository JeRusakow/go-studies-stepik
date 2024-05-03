/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
Функция work() ничего не принимает и не возвращает.
*/

package main

import (
	"sync"
	"time"
)

func work() {
	time.Sleep(time.Second)
	println("Work done")
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			work()
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}
